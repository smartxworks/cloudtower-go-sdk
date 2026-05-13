package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	goruntime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

var (
	ErrActivePassiveNoEndpoints      = errors.New("active-passive client requires at least one endpoint")
	ErrActivePassiveDuplicateHost    = errors.New("active-passive client endpoints must be unique")
	ErrActivePassiveNoActiveHost     = errors.New("active-passive discover found no active host")
	ErrActivePassiveMultipleActives  = errors.New("active-passive discover found multiple active hosts")
	ErrActivePassiveRetryExhausted   = errors.New("active-passive request retry exhausted after discover")
	ErrActivePassiveMissingLocation  = errors.New("active-passive redirect is missing location header")
	ErrActivePassiveFailoverRequired = errors.New("active-passive failover required")
	ErrActivePassiveUnknownHost      = errors.New("active-passive current active host is not configured")
	ErrActivePassiveUnsupportedAuth  = errors.New("client transport does not support default authentication")
	errActivePassiveNoResponseReader = errors.New("client operation does not provide a response reader")
)

// FailoverStrategy controls when active-passive client discovers and retries endpoints.
type FailoverStrategy int

const (
	FailoverStrategyDefault FailoverStrategy = iota
	FailoverStrategyManualFailover
	FailoverStrategyAlwaysProbe
)

// ActivePassiveClientConfig contains transport-related options for the active-passive client.
type ActivePassiveClientConfig struct {
	Endpoints        []string
	BasePath         string
	Schemes          []string
	UserConfig       *UserConfig
	ProbeTimeout     time.Duration
	HTTPClient       *http.Client
	FailoverStrategy FailoverStrategy
	formats          *strfmt.Registry
}

// DefaultActivePassiveClientConfig returns the default config for an active-passive client.
func DefaultActivePassiveClientConfig() *ActivePassiveClientConfig {
	return &ActivePassiveClientConfig{
		BasePath:     DefaultBasePath,
		Schemes:      DefaultSchemes,
		ProbeTimeout: httptransport.DefaultTimeout,
	}
}

// ActivePassiveTransport routes requests to the current active endpoint.
type ActivePassiveTransport struct {
	mu                sync.RWMutex
	endpoints         map[string]*activePassiveEndpoint
	orderedHosts      []string
	currentActiveHost string
	discoverWait      chan struct{}
	discoverErr       error
	probeTimeout      time.Duration
	failoverStrategy  FailoverStrategy
}

// ActivePassiveClient wraps Cloudtower with active-passive specific state accessors.
type ActivePassiveClient struct {
	*Cloudtower
	transport *ActivePassiveTransport
}

type activePassiveEndpoint struct {
	host         string
	runtime      *httptransport.Runtime
	httpClient   *http.Client
	probeRuntime *httptransport.Runtime
	probeClient  *http.Client
}

type activePassiveSubmitState int

const (
	activePassiveSubmitSuccess activePassiveSubmitState = iota
	activePassiveSubmitRecognizableError
	activePassiveSubmitUnrecognizableError
	activePassiveSubmitSwitchSignal
	activePassiveSubmitLocalError
)

type activePassiveSubmitResult struct {
	state  activePassiveSubmitState
	result interface{}
	err    error
}

type activePassiveProbeState int

const (
	activePassiveProbeActive activePassiveProbeState = iota
	activePassiveProbePassive
)

type activePassiveHTTPResponse struct {
	response *http.Response
}

func (r activePassiveHTTPResponse) Code() int {
	return r.response.StatusCode
}

func (r activePassiveHTTPResponse) Message() string {
	return r.response.Status
}

func (r activePassiveHTTPResponse) GetHeader(name string) string {
	return r.response.Header.Get(name)
}

func (r activePassiveHTTPResponse) GetHeaders(name string) []string {
	return r.response.Header.Values(name)
}

func (r activePassiveHTTPResponse) Body() io.ReadCloser {
	return r.response.Body
}

// NewActivePassiveClient creates a new active-passive cloudtower client.
func NewActivePassiveClient(ctx context.Context, clientConfig ActivePassiveClientConfig) (*ActivePassiveClient, error) {
	if ctx == nil {
		timeout := clientConfig.ProbeTimeout
		if timeout == 0 {
			timeout = httptransport.DefaultTimeout
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
		defer cancel()
	}

	formats := strfmt.Default
	if clientConfig.formats != nil {
		formats = *clientConfig.formats
	}

	transport, err := newActivePassiveTransport(clientConfig, formats)
	if err != nil {
		return nil, err
	}

	cloudtowerClient := New(transport, formats)
	client := &ActivePassiveClient{
		Cloudtower: cloudtowerClient,
		transport:  transport,
	}
	if clientConfig.UserConfig != nil {
		host, err := transport.ensureActiveHost(ctx)
		if err != nil {
			return nil, err
		}

		if err := loginWithUserConfig(client.Cloudtower, host, *clientConfig.UserConfig); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// NewActivePassiveWithUserConfig creates a new active-passive client and logs in.
func NewActivePassiveWithUserConfig(ctx context.Context, clientConfig ActivePassiveClientConfig, userConfig UserConfig) (*ActivePassiveClient, error) {
	clientConfig.UserConfig = &userConfig
	return NewActivePassiveClient(ctx, clientConfig)
}

func newActivePassiveTransport(clientConfig ActivePassiveClientConfig, formats strfmt.Registry) (*ActivePassiveTransport, error) {
	cfg := clientConfig
	if cfg.BasePath == "" {
		cfg.BasePath = DefaultBasePath
	}
	if len(cfg.Schemes) == 0 {
		cfg.Schemes = DefaultSchemes
	}
	if cfg.ProbeTimeout == 0 {
		cfg.ProbeTimeout = httptransport.DefaultTimeout
	}
	if len(cfg.Endpoints) == 0 {
		return nil, ErrActivePassiveNoEndpoints
	}

	transport := &ActivePassiveTransport{
		endpoints:        make(map[string]*activePassiveEndpoint, len(cfg.Endpoints)),
		orderedHosts:     make([]string, 0, len(cfg.Endpoints)),
		probeTimeout:     cfg.ProbeTimeout,
		failoverStrategy: cfg.FailoverStrategy,
	}

	for _, endpoint := range cfg.Endpoints {
		host := strings.TrimSpace(endpoint)
		if host == "" {
			return nil, ErrActivePassiveNoEndpoints
		}
		if _, exists := transport.endpoints[host]; exists {
			return nil, fmt.Errorf("%w: %s", ErrActivePassiveDuplicateHost, host)
		}

		httpClient := cloneNoRedirectHTTPClient(cfg.HTTPClient)
		runtime := httptransport.NewWithClient(host, cfg.BasePath, cfg.Schemes, httpClient)
		runtime.Formats = formats
		probeRuntime := httptransport.NewWithClient(host, "/", cfg.Schemes, httpClient)
		probeRuntime.Formats = formats

		transport.endpoints[host] = &activePassiveEndpoint{
			host:         host,
			runtime:      runtime,
			httpClient:   httpClient,
			probeRuntime: probeRuntime,
			probeClient:  httpClient,
		}
		transport.orderedHosts = append(transport.orderedHosts, host)
	}

	return transport, nil
}

// Submit routes the client operation to the current active host.
func (t *ActivePassiveTransport) Submit(op *goruntime.ClientOperation) (interface{}, error) {
	host, err := t.hostForRequest(op.Context)
	if err != nil {
		return nil, err
	}

	result := t.submitToHost(host, op)
	switch result.state {
	case activePassiveSubmitSuccess:
		return result.result, nil
	case activePassiveSubmitRecognizableError, activePassiveSubmitLocalError:
		return nil, result.err
	case activePassiveSubmitUnrecognizableError:
		t.clearCurrentActiveHostIf(host)
		return nil, result.err
	case activePassiveSubmitSwitchSignal:
		t.clearCurrentActiveHostIf(host)
		if t.failoverStrategy == FailoverStrategyManualFailover || t.failoverStrategy == FailoverStrategyAlwaysProbe {
			return nil, fmt.Errorf("%w: %v", ErrActivePassiveFailoverRequired, result.err)
		}

		nextHost, err := t.ensureActiveHost(op.Context)
		if err != nil {
			return nil, err
		}

		retry := t.submitToHost(nextHost, op)
		switch retry.state {
		case activePassiveSubmitSuccess:
			return retry.result, nil
		case activePassiveSubmitRecognizableError, activePassiveSubmitLocalError:
			return nil, retry.err
		case activePassiveSubmitUnrecognizableError:
			t.clearCurrentActiveHostIf(nextHost)
			return nil, retry.err
		case activePassiveSubmitSwitchSignal:
			t.clearCurrentActiveHostIf(nextHost)
			return nil, ErrActivePassiveRetryExhausted
		default:
			return nil, retry.err
		}
	default:
		return nil, result.err
	}
}

func (t *ActivePassiveTransport) hostForRequest(ctx context.Context) (string, error) {
	if t.failoverStrategy == FailoverStrategyAlwaysProbe {
		t.clearCurrentActiveHost()
	}

	return t.ensureActiveHost(ctx)
}

// SetDefaultAuthentication updates the auth info writer on every endpoint runtime.
func (t *ActivePassiveTransport) SetDefaultAuthentication(auth goruntime.ClientAuthInfoWriter) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, endpoint := range t.endpoints {
		endpoint.runtime.DefaultAuthentication = auth
	}
}

// CurrentActiveHost returns the host currently used for request routing.
func (t *ActivePassiveTransport) CurrentActiveHost() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.currentActiveHost
}

// SetDefaultAuthentication updates the authentication writer for both regular and active-passive clients.
func (c *Cloudtower) SetDefaultAuthentication(auth goruntime.ClientAuthInfoWriter) {
	switch transport := c.Transport.(type) {
	case *ActivePassiveTransport:
		transport.SetDefaultAuthentication(auth)
	case *httptransport.Runtime:
		transport.DefaultAuthentication = auth
	}
}

// CurrentActiveHost returns the last discovered active host.
func (c *ActivePassiveClient) CurrentActiveHost() string {
	if c == nil || c.transport == nil {
		return ""
	}
	return c.transport.CurrentActiveHost()
}

func (t *ActivePassiveTransport) ensureActiveHost(ctx context.Context) (string, error) {
	if host := t.CurrentActiveHost(); host != "" {
		return host, nil
	}

	for {
		t.mu.Lock()
		if t.currentActiveHost != "" {
			host := t.currentActiveHost
			t.mu.Unlock()
			return host, nil
		}

		if wait := t.discoverWait; wait != nil {
			t.mu.Unlock()
			if err := waitForDiscover(ctx, wait); err != nil {
				return "", err
			}

			t.mu.RLock()
			host := t.currentActiveHost
			err := t.discoverErr
			t.mu.RUnlock()
			if host != "" {
				return host, nil
			}
			if err != nil {
				return "", err
			}
			continue
		}

		wait := make(chan struct{})
		t.discoverWait = wait
		t.mu.Unlock()

		host, err := t.discover(ctx)

		t.mu.Lock()
		if err == nil {
			t.currentActiveHost = host
		}
		t.discoverErr = err
		close(wait)
		t.discoverWait = nil
		host = t.currentActiveHost
		t.mu.Unlock()

		if err != nil {
			return "", err
		}
		return host, nil
	}
}

func (t *ActivePassiveTransport) discover(ctx context.Context) (string, error) {
	activeHosts := make([]string, 0, 1)
	failures := make([]string, 0)

	for _, host := range t.orderedHosts {
		state, err := t.probeHost(ctx, host)
		if err != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", host, err))
			continue
		}
		if state == activePassiveProbeActive {
			activeHosts = append(activeHosts, host)
		}
	}

	switch len(activeHosts) {
	case 1:
		return activeHosts[0], nil
	case 0:
		if len(failures) == 0 {
			return "", ErrActivePassiveNoActiveHost
		}
		return "", fmt.Errorf("%w: %s", ErrActivePassiveNoActiveHost, strings.Join(failures, "; "))
	default:
		return "", fmt.Errorf("%w: %s", ErrActivePassiveMultipleActives, strings.Join(activeHosts, ", "))
	}
}

func (t *ActivePassiveTransport) probeHost(ctx context.Context, host string) (activePassiveProbeState, error) {
	endpoint, ok := t.endpoints[host]
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrActivePassiveUnknownHost, host)
	}

	reqCtx, cancel := context.WithTimeout(defaultContext(ctx), t.probeTimeout)
	defer cancel()

	req, err := endpoint.probeRuntime.CreateHttpRequest(activePassiveProbeOperation(reqCtx))
	if err != nil {
		return 0, err
	}

	resp, err := endpoint.probeClient.Do(req.WithContext(reqCtx))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return activePassiveProbeActive, nil
	case http.StatusTemporaryRedirect:
		return activePassiveProbePassive, nil
	default:
		return 0, goruntime.NewAPIError("probe active-passive returned unexpected status", activePassiveHTTPResponse{response: resp}, resp.StatusCode)
	}
}

func activePassiveProbeOperation(ctx context.Context) *goruntime.ClientOperation {
	return &goruntime.ClientOperation{
		ID:                 "probe-active-passive",
		Method:             http.MethodGet,
		PathPattern:        "/api/healthz",
		ProducesMediaTypes: []string{goruntime.JSONMime},
		Params: goruntime.ClientRequestWriterFunc(func(goruntime.ClientRequest, strfmt.Registry) error {
			return nil
		}),
		Reader: goruntime.ClientResponseReaderFunc(func(goruntime.ClientResponse, goruntime.Consumer) (interface{}, error) {
			return nil, nil
		}),
		Context: ctx,
	}
}

func (t *ActivePassiveTransport) submitToHost(host string, op *goruntime.ClientOperation) activePassiveSubmitResult {
	endpoint, ok := t.endpoints[host]
	if !ok {
		return activePassiveSubmitResult{
			state: activePassiveSubmitLocalError,
			err:   fmt.Errorf("%w: %s", ErrActivePassiveUnknownHost, host),
		}
	}

	req, err := endpoint.runtime.CreateHttpRequest(op)
	if err != nil {
		return activePassiveSubmitResult{
			state: activePassiveSubmitLocalError,
			err:   err,
		}
	}

	ctx, cancel := requestContext(endpoint.runtime.Context, op.Context, operationTimeout(op.Params))
	defer cancel()

	resp, err := effectiveHTTPClient(op.Client, endpoint.httpClient).Do(req.WithContext(ctx))
	if err != nil {
		return activePassiveSubmitResult{
			state: activePassiveSubmitUnrecognizableError,
			err:   err,
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTemporaryRedirect {
		return activePassiveSubmitResult{
			state: activePassiveSubmitSwitchSignal,
			err:   goruntime.NewAPIError("active-passive switch signal", activePassiveHTTPResponse{response: resp}, resp.StatusCode),
		}
	}

	if op.Reader == nil {
		return activePassiveSubmitResult{
			state: activePassiveSubmitLocalError,
			err:   errActivePassiveNoResponseReader,
		}
	}

	consumer, err := responseConsumer(endpoint.runtime, resp)
	if err != nil {
		return activePassiveSubmitResult{
			state: activePassiveSubmitUnrecognizableError,
			err:   err,
		}
	}

	result, err := op.Reader.ReadResponse(activePassiveHTTPResponse{response: resp}, consumer)
	if err == nil {
		return activePassiveSubmitResult{
			state:  activePassiveSubmitSuccess,
			result: result,
		}
	}

	if isUnrecognizableResponseError(err) {
		return activePassiveSubmitResult{
			state: activePassiveSubmitUnrecognizableError,
			err:   err,
		}
	}

	return activePassiveSubmitResult{
		state: activePassiveSubmitRecognizableError,
		err:   err,
	}
}

func (t *ActivePassiveTransport) clearCurrentActiveHostIf(host string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.currentActiveHost == host {
		t.currentActiveHost = ""
	}
}

func (t *ActivePassiveTransport) clearCurrentActiveHost() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.currentActiveHost = ""
}

func cloneNoRedirectHTTPClient(base *http.Client) *http.Client {
	if base == nil {
		return &http.Client{CheckRedirect: stopRedirect}
	}

	cloned := *base
	cloned.CheckRedirect = stopRedirect
	return &cloned
}

func stopRedirect(*http.Request, []*http.Request) error {
	return http.ErrUseLastResponse
}

func effectiveHTTPClient(opClient *http.Client, fallback *http.Client) *http.Client {
	if opClient == nil {
		return fallback
	}
	return cloneNoRedirectHTTPClient(opClient)
}

func requestContext(runtimeCtx, opCtx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	if opCtx != nil {
		return context.WithCancel(opCtx)
	}

	parent := runtimeCtx
	if parent == nil {
		parent = context.Background()
	}

	if timeout <= 0 {
		return context.WithCancel(parent)
	}

	return context.WithTimeout(parent, timeout)
}

func defaultContext(ctx context.Context) context.Context {
	if ctx != nil {
		return ctx
	}
	return context.Background()
}

func operationTimeout(params goruntime.ClientRequestWriter) time.Duration {
	if params == nil {
		return httptransport.DefaultTimeout
	}

	value := reflect.ValueOf(params)
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return httptransport.DefaultTimeout
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return httptransport.DefaultTimeout
	}

	timeoutField := value.FieldByName("timeout")
	if !timeoutField.IsValid() || timeoutField.Kind() != reflect.Int64 {
		return httptransport.DefaultTimeout
	}

	timeout := time.Duration(timeoutField.Int())
	if timeout <= 0 {
		return httptransport.DefaultTimeout
	}

	return timeout
}

func responseConsumer(runtime *httptransport.Runtime, resp *http.Response) (goruntime.Consumer, error) {
	contentType := resp.Header.Get(goruntime.HeaderContentType)
	if contentType == "" {
		contentType = runtime.DefaultMediaType
	}

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("parse content type: %w", err)
	}

	consumer, ok := runtime.Consumers[mediaType]
	if ok {
		return consumer, nil
	}

	consumer, ok = runtime.Consumers[goruntime.DefaultMime]
	if ok {
		return consumer, nil
	}

	return nil, fmt.Errorf("no consumer for %q", contentType)
}

func isUnrecognizableResponseError(err error) bool {
	var apiErr *goruntime.APIError
	if errors.As(err, &apiErr) {
		return true
	}

	var unexpectedErr *models.UnexpectedError
	return errors.As(err, &unexpectedErr)
}

func waitForDiscover(ctx context.Context, wait chan struct{}) error {
	if ctx == nil {
		<-wait
		return nil
	}

	select {
	case <-wait:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
