package client

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	goruntime "github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

var errActivePassiveTestRecognizable = errors.New("recognizable api error")

func newTestActivePassiveTransport(t *testing.T, serverURLs ...string) *ActivePassiveTransport {
	t.Helper()

	endpoints := make([]ActivePassiveEndpointConfig, 0, len(serverURLs))
	for _, serverURL := range serverURLs {
		parsedURL := parseTestServerURL(t, serverURL)
		endpoints = append(endpoints, ActivePassiveEndpointConfig{
			Host:     parsedURL.Host,
			Schemes:  []string{parsedURL.Scheme},
			BasePath: "/",
		})
	}

	transport, err := newActivePassiveTransport(ActivePassiveClientConfig{
		Endpoints:    endpoints,
		ProbeTimeout: time.Second,
	}, strfmt.Default)
	if err != nil {
		t.Fatalf("newActivePassiveTransport() error = %v", err)
	}

	return transport
}

func TestActivePassiveProbeUsesConfiguredHTTPClientTLS(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	transport, err := newActivePassiveTransport(ActivePassiveClientConfig{
		Endpoints: []ActivePassiveEndpointConfig{
			{Host: parseTestServerURL(t, server.URL).Host, Schemes: []string{parseTestServerURL(t, server.URL).Scheme}, BasePath: "/v2/api"},
		},
		ProbeTimeout: time.Second,
		HTTPClient:   server.Client(),
	}, strfmt.Default)
	if err != nil {
		t.Fatalf("newActivePassiveTransport() error = %v", err)
	}

	activeEndpoint, err := transport.ensureActiveEndpoint(context.Background())
	if err != nil {
		t.Fatalf("ensureActiveEndpoint() error = %v", err)
	}
	if activeEndpoint != server.URL {
		t.Fatalf("ensureActiveEndpoint() = %q, want %q", activeEndpoint, server.URL)
	}
}

func TestActivePassiveEndpointPathPrefixesProbeAndAPIRequests(t *testing.T) {
	var sawProbe bool
	var sawRequest bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/gateway/custom-api/api/healthz":
			sawProbe = true
			w.WriteHeader(http.StatusOK)
		case "/gateway/custom-api/v2/api/test":
			sawRequest = true
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("active"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	transport, err := newActivePassiveTransport(ActivePassiveClientConfig{
		Endpoints: []ActivePassiveEndpointConfig{
			{
				Host:           parseTestServerURL(t, server.URL).Host,
				Schemes:        []string{parseTestServerURL(t, server.URL).Scheme},
				CommonBasePath: "/gateway/custom-api",
				BasePath:       "/v2/api",
			},
		},
		ProbeTimeout: time.Second,
	}, strfmt.Default)
	if err != nil {
		t.Fatalf("newActivePassiveTransport() error = %v", err)
	}

	result, err := transport.Submit(testActivePassiveOperation())
	if err != nil {
		t.Fatalf("Submit() error = %v", err)
	}
	if result != "active" {
		t.Fatalf("Submit() result = %v, want active", result)
	}
	if !sawProbe {
		t.Fatal("probe did not use endpoint path prefix")
	}
	if !sawRequest {
		t.Fatal("request did not use endpoint path prefix plus API base path")
	}
}

func TestActivePassiveEndpointUsesDefaultPaths(t *testing.T) {
	var sawProbe bool
	var sawRequest bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/tower-a/api/healthz":
			sawProbe = true
			w.WriteHeader(http.StatusOK)
		case "/tower-a/test":
			sawRequest = true
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("active"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	transport, err := newActivePassiveTransport(ActivePassiveClientConfig{
		Endpoints: []ActivePassiveEndpointConfig{
			{
				Host:           parseTestServerURL(t, server.URL).Host,
				Schemes:        []string{parseTestServerURL(t, server.URL).Scheme},
				CommonBasePath: "/tower-a",
			},
		},
		ProbeTimeout: time.Second,
	}, strfmt.Default)
	if err != nil {
		t.Fatalf("newActivePassiveTransport() error = %v", err)
	}

	result, err := transport.Submit(testActivePassiveOperation())
	if err != nil {
		t.Fatalf("Submit() error = %v", err)
	}
	if result != "active" {
		t.Fatalf("Submit() result = %v, want active", result)
	}
	if !sawProbe {
		t.Fatal("probe did not use default probe path")
	}
	if !sawRequest {
		t.Fatal("request did not use default API path")
	}
}

func TestActivePassiveAppliesCommonBasePathAcrossHosts(t *testing.T) {
	passive := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/tower-b/api/healthz":
			w.WriteHeader(http.StatusTemporaryRedirect)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer passive.Close()

	active := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/tower-b/api/healthz":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer active.Close()

	transport, err := newActivePassiveTransport(ActivePassiveClientConfig{
		Endpoints: []ActivePassiveEndpointConfig{
			{Host: parseTestServerURL(t, passive.URL).Host, Schemes: []string{parseTestServerURL(t, passive.URL).Scheme}, CommonBasePath: "/tower-b", BasePath: "/v2/api"},
			{Host: parseTestServerURL(t, active.URL).Host, Schemes: []string{parseTestServerURL(t, active.URL).Scheme}, CommonBasePath: "/tower-b", BasePath: "/v2/api"},
		},
		ProbeTimeout: time.Second,
	}, strfmt.Default)
	if err != nil {
		t.Fatalf("newActivePassiveTransport() error = %v", err)
	}

	activeEndpoint, err := transport.ensureActiveEndpoint(context.Background())
	if err != nil {
		t.Fatalf("ensureActiveEndpoint() error = %v", err)
	}
	if want := active.URL + "/tower-b"; activeEndpoint != want {
		t.Fatalf("ensureActiveEndpoint() = %q, want %q", activeEndpoint, want)
	}

	for _, endpoint := range transport.endpoints {
		if endpoint.runtime.BasePath != "/tower-b/v2/api" {
			t.Fatalf("endpoint runtime BasePath = %q, want /tower-b/v2/api", endpoint.runtime.BasePath)
		}
		if endpoint.probeRuntime.BasePath != "/tower-b/api/healthz" {
			t.Fatalf("endpoint probe runtime BasePath = %q, want /tower-b/api/healthz", endpoint.probeRuntime.BasePath)
		}
	}
}

func newTestActivePassiveTransportWithStrategy(t *testing.T, strategy FailoverStrategy, serverURLs ...string) *ActivePassiveTransport {
	t.Helper()

	endpoints := make([]ActivePassiveEndpointConfig, 0, len(serverURLs))
	for _, serverURL := range serverURLs {
		parsedURL := parseTestServerURL(t, serverURL)
		endpoints = append(endpoints, ActivePassiveEndpointConfig{
			Host:     parsedURL.Host,
			Schemes:  []string{parsedURL.Scheme},
			BasePath: "/",
		})
	}

	transport, err := newActivePassiveTransport(ActivePassiveClientConfig{
		Endpoints:        endpoints,
		ProbeTimeout:     time.Second,
		FailoverStrategy: strategy,
	}, strfmt.Default)
	if err != nil {
		t.Fatalf("newActivePassiveTransport() error = %v", err)
	}

	return transport
}

func testActivePassiveOperation() *goruntime.ClientOperation {
	return &goruntime.ClientOperation{
		ID:                 "test-active-passive",
		Method:             http.MethodGet,
		PathPattern:        "/test",
		ProducesMediaTypes: []string{goruntime.JSONMime},
		Params: goruntime.ClientRequestWriterFunc(func(goruntime.ClientRequest, strfmt.Registry) error {
			return nil
		}),
		Reader: goruntime.ClientResponseReaderFunc(func(response goruntime.ClientResponse, consumer goruntime.Consumer) (interface{}, error) {
			switch response.Code() {
			case http.StatusOK:
				body, err := ioutil.ReadAll(response.Body())
				if err != nil {
					return nil, err
				}
				return string(body), nil
			case http.StatusBadRequest:
				return nil, errActivePassiveTestRecognizable
			default:
				return nil, goruntime.NewAPIError("test active-passive", response, response.Code())
			}
		}),
		Context: context.Background(),
	}
}

func setTestCurrentActiveEndpoint(transport *ActivePassiveTransport, endpointBaseURL string) {
	transport.mu.Lock()
	defer transport.mu.Unlock()
	transport.currentActiveEndpoint = endpointBaseURL
}

func parseTestServerURL(t *testing.T, rawURL string) *url.URL {
	t.Helper()
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("parse test server URL: %v", err)
	}
	return parsedURL
}
