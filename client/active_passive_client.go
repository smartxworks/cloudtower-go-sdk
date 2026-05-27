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
	ErrActivePassiveNoEndpoints       = errors.New("active-passive client requires at least one endpoint")
	ErrActivePassiveInvalidEndpoint   = errors.New("active-passive client endpoint must be an absolute URL without query, fragment, or user info")
	ErrActivePassiveDuplicateEndpoint = errors.New("active-passive client endpoints must be unique")
	ErrActivePassiveNoActiveHost      = errors.New("active-passive discover found no active host")
	ErrActivePassiveMultipleActives   = errors.New("active-passive discover found multiple active hosts")
	ErrActivePassiveRetryExhausted    = errors.New("active-passive request retry exhausted after discover")
	ErrActivePassiveMissingLocation   = errors.New("active-passive redirect is missing location header")
	ErrActivePassiveFailoverRequired  = errors.New("active-passive failover required")
	ErrActivePassiveUnknownEndpoint   = errors.New("active-passive current active endpoint is not configured")
	ErrActivePassiveUnsupportedAuth   = errors.New("client transport does not support default authentication")
	errActivePassiveNoResponseReader  = errors.New("client operation does not provide a response reader")
)

// FailoverStrategy controls when active-passive client discovers and retries endpoints.
type FailoverStrategy int

const (
	FailoverStrategyDefault FailoverStrategy = iota
	FailoverStrategyManualFailover
	FailoverStrategyAlwaysProbe
)

const (
	activePassiveDefaultProbePath = "/api/healthz"
)

// ActivePassiveClientConfig contains transport-related options for the active-passive client.
type ActivePassiveClientConfig struct {
	// Endpoints contains candidate active/passive endpoints.
	Endpoints        []ActivePassiveEndpointConfig
	UserConfig       *UserConfig
	ProbeTimeout     time.Duration
	HTTPClient       *http.Client
	FailoverStrategy FailoverStrategy
	formats          *strfmt.Registry
}

// ActivePassiveEndpointConfig contains one active/passive endpoint and its path configuration.
type ActivePassiveEndpointConfig struct {
	Host           string
	Schemes        []string
	CommonBasePath string
	BasePath       string
	ProbePath      string
}

// ActivePassiveTransport routes requests to the current active endpoint.
type ActivePassiveTransport struct {
	mu                    sync.RWMutex
	endpoints             map[string]*activePassiveEndpoint
	currentActiveEndpoint string
	discoverWait          chan struct{}
	discoverErr           error
	probeTimeout          time.Duration
	failoverStrategy      FailoverStrategy
}

// ActivePassiveClient wraps Cloudtower with active-passive specific state accessors.
type ActivePassiveClient struct {
	*Cloudtower
	transport *ActivePassiveTransport
}

type activePassiveEndpoint struct {
	baseURL      string
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
		endpointBaseURL, err := transport.ensureActiveEndpoint(ctx)
		if err != nil {
			return nil, err
		}

		if err := loginWithUserConfig(client.Cloudtower, endpointBaseURL, *clientConfig.UserConfig); err != nil {
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
	if cfg.ProbeTimeout == 0 {
		cfg.ProbeTimeout = httptransport.DefaultTimeout
	}
	if len(cfg.Endpoints) == 0 {
		return nil, ErrActivePassiveNoEndpoints
	}

	transport := &ActivePassiveTransport{
		endpoints:        make(map[string]*activePassiveEndpoint, len(cfg.Endpoints)),
		probeTimeout:     cfg.ProbeTimeout,
		failoverStrategy: cfg.FailoverStrategy,
	}

	for _, endpointConfig := range cfg.Endpoints {
		if len(endpointConfig.Schemes) == 0 {
			endpointConfig.Schemes = DefaultSchemes
		}
		if endpointConfig.BasePath == "" {
			endpointConfig.BasePath = DefaultBasePath
		}
		if endpointConfig.ProbePath == "" {
			endpointConfig.ProbePath = activePassiveDefaultProbePath
		}

		endpoint := newActivePassiveEndpoint(endpointConfig, cfg.HTTPClient, formats)
		if _, exists := transport.endpoints[endpoint.baseURL]; exists {
			return nil, fmt.Errorf("%w: %s", ErrActivePassiveDuplicateEndpoint, endpoint.baseURL)
		}

		transport.endpoints[endpoint.baseURL] = endpoint
	}

	return transport, nil
}

func newActivePassiveEndpoint(endpointConfig ActivePassiveEndpointConfig, baseHTTPClient *http.Client, formats strfmt.Registry) *activePassiveEndpoint {
	httpClient := cloneNoRedirectHTTPClient(baseHTTPClient)
	host := strings.TrimSpace(endpointConfig.Host)
	commonBasePath := normalizeURLPath(endpointConfig.CommonBasePath)
	apiBasePath := joinURLPaths(commonBasePath, endpointConfig.BasePath)
	probePath := joinURLPaths(commonBasePath, endpointConfig.ProbePath)
	baseURL := strings.TrimRight(endpointConfig.Schemes[0]+"://"+host+commonBasePath, "/")

	runtime := httptransport.NewWithClient(host, apiBasePath, endpointConfig.Schemes, httpClient)
	runtime.Formats = formats
	probeRuntime := httptransport.NewWithClient(host, probePath, endpointConfig.Schemes, httpClient)
	probeRuntime.Formats = formats

	return &activePassiveEndpoint{
		baseURL:      baseURL,
		runtime:      runtime,
		httpClient:   httpClient,
		probeRuntime: probeRuntime,
		probeClient:  httpClient,
	}
}

func normalizeURLPath(value string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" || trimmed == "/" {
		return ""
	}
	if !strings.HasPrefix(trimmed, "/") {
		trimmed = "/" + trimmed
	}
	return strings.TrimRight(trimmed, "/")
}

func joinURLPaths(parts ...string) string {
	joined := ""
	for _, part := range parts {
		normalized := normalizeURLPath(part)
		if normalized == "" {
			continue
		}
		joined += normalized
	}
	if joined == "" {
		return "/"
	}
	return joined
}

// Submit routes the client operation to the current active endpoint.
func (t *ActivePassiveTransport) Submit(op *goruntime.ClientOperation) (interface{}, error) {
	endpointBaseURL, err := t.endpointForRequest(op.Context)
	if err != nil {
		return nil, err
	}

	result := t.submitToEndpoint(endpointBaseURL, op)
	switch result.state {
	case activePassiveSubmitSuccess:
		return result.result, nil
	case activePassiveSubmitRecognizableError, activePassiveSubmitLocalError:
		return nil, result.err
	case activePassiveSubmitUnrecognizableError:
		t.clearCurrentActiveEndpointIf(endpointBaseURL)
		return nil, result.err
	case activePassiveSubmitSwitchSignal:
		t.clearCurrentActiveEndpointIf(endpointBaseURL)
		if t.failoverStrategy == FailoverStrategyManualFailover || t.failoverStrategy == FailoverStrategyAlwaysProbe {
			return nil, fmt.Errorf("%w: %v", ErrActivePassiveFailoverRequired, result.err)
		}

		nextEndpointBaseURL, err := t.ensureActiveEndpoint(op.Context)
		if err != nil {
			return nil, err
		}

		retry := t.submitToEndpoint(nextEndpointBaseURL, op)
		switch retry.state {
		case activePassiveSubmitSuccess:
			return retry.result, nil
		case activePassiveSubmitRecognizableError, activePassiveSubmitLocalError:
			return nil, retry.err
		case activePassiveSubmitUnrecognizableError:
			t.clearCurrentActiveEndpointIf(nextEndpointBaseURL)
			return nil, retry.err
		case activePassiveSubmitSwitchSignal:
			t.clearCurrentActiveEndpointIf(nextEndpointBaseURL)
			return nil, ErrActivePassiveRetryExhausted
		default:
			return nil, retry.err
		}
	default:
		return nil, result.err
	}
}

func (t *ActivePassiveTransport) endpointForRequest(ctx context.Context) (string, error) {
	if t.failoverStrategy == FailoverStrategyAlwaysProbe {
		t.clearCurrentActiveEndpoint()
	}

	return t.ensureActiveEndpoint(ctx)
}

// SetDefaultAuthentication updates the auth info writer on every endpoint runtime.
func (t *ActivePassiveTransport) SetDefaultAuthentication(auth goruntime.ClientAuthInfoWriter) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, endpoint := range t.endpoints {
		endpoint.runtime.DefaultAuthentication = auth
	}
}

// CurrentActiveEndpoint returns the endpoint base URL currently used for request routing.
func (t *ActivePassiveTransport) CurrentActiveEndpoint() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.currentActiveEndpoint
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

// CurrentActiveEndpoint returns the last discovered active endpoint base URL.
func (c *ActivePassiveClient) CurrentActiveEndpoint() string {
	if c == nil || c.transport == nil {
		return ""
	}
	return c.transport.CurrentActiveEndpoint()
}

func (t *ActivePassiveTransport) ensureActiveEndpoint(ctx context.Context) (string, error) {
	if endpointBaseURL := t.CurrentActiveEndpoint(); endpointBaseURL != "" {
		return endpointBaseURL, nil
	}

	for {
		t.mu.Lock()
		if t.currentActiveEndpoint != "" {
			endpointBaseURL := t.currentActiveEndpoint
			t.mu.Unlock()
			return endpointBaseURL, nil
		}

		if wait := t.discoverWait; wait != nil {
			t.mu.Unlock()
			if err := waitForDiscover(ctx, wait); err != nil {
				return "", err
			}

			t.mu.RLock()
			endpointBaseURL := t.currentActiveEndpoint
			err := t.discoverErr
			t.mu.RUnlock()
			if endpointBaseURL != "" {
				return endpointBaseURL, nil
			}
			if err != nil {
				return "", err
			}
			continue
		}

		wait := make(chan struct{})
		t.discoverWait = wait
		t.mu.Unlock()

		endpointBaseURL, err := t.discover(ctx)

		t.mu.Lock()
		if err == nil {
			t.currentActiveEndpoint = endpointBaseURL
		}
		t.discoverErr = err
		close(wait)
		t.discoverWait = nil
		endpointBaseURL = t.currentActiveEndpoint
		t.mu.Unlock()

		if err != nil {
			return "", err
		}
		return endpointBaseURL, nil
	}
}

func (t *ActivePassiveTransport) discover(ctx context.Context) (string, error) {
	activeEndpoints := make([]string, 0, 1)
	failures := make([]string, 0)

	for _, endpoint := range t.endpoints {
		state, err := t.probeEndpoint(ctx, endpoint)
		if err != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", endpoint.baseURL, err))
			continue
		}
		if state == activePassiveProbeActive {
			activeEndpoints = append(activeEndpoints, endpoint.baseURL)
		}
	}

	switch len(activeEndpoints) {
	case 1:
		return activeEndpoints[0], nil
	case 0:
		if len(failures) == 0 {
			return "", ErrActivePassiveNoActiveHost
		}
		return "", fmt.Errorf("%w: %s", ErrActivePassiveNoActiveHost, strings.Join(failures, "; "))
	default:
		return "", fmt.Errorf("%w: %s", ErrActivePassiveMultipleActives, strings.Join(activeEndpoints, ", "))
	}
}

func (t *ActivePassiveTransport) probeEndpoint(ctx context.Context, endpoint *activePassiveEndpoint) (activePassiveProbeState, error) {
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
		PathPattern:        "/",
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

func (t *ActivePassiveTransport) submitToEndpoint(endpointBaseURL string, op *goruntime.ClientOperation) activePassiveSubmitResult {
	endpoint, ok := t.endpoints[endpointBaseURL]
	if !ok {
		return activePassiveSubmitResult{
			state: activePassiveSubmitLocalError,
			err:   fmt.Errorf("%w: %s", ErrActivePassiveUnknownEndpoint, endpointBaseURL),
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

func (t *ActivePassiveTransport) clearCurrentActiveEndpointIf(endpointBaseURL string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.currentActiveEndpoint == endpointBaseURL {
		t.currentActiveEndpoint = ""
	}
}

func (t *ActivePassiveTransport) clearCurrentActiveEndpoint() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.currentActiveEndpoint = ""
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
