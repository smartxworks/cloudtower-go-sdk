package client

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	goruntime "github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

var errActivePassiveTestRecognizable = errors.New("recognizable api error")

func newTestActivePassiveTransport(t *testing.T, serverURLs ...string) *ActivePassiveTransport {
	t.Helper()

	endpoints := make([]string, 0, len(serverURLs))
	for _, serverURL := range serverURLs {
		endpoints = append(endpoints, testServerHost(t, serverURL))
	}

	transport, err := newActivePassiveTransport(ActivePassiveClientConfig{
		Endpoints:    endpoints,
		BasePath:     "/",
		Schemes:      []string{"http"},
		ProbeTimeout: time.Second,
	}, strfmt.Default)
	if err != nil {
		t.Fatalf("newActivePassiveTransport() error = %v", err)
	}

	return transport
}

func newTestActivePassiveTransportWithStrategy(t *testing.T, strategy FailoverStrategy, serverURLs ...string) *ActivePassiveTransport {
	t.Helper()

	endpoints := make([]string, 0, len(serverURLs))
	for _, serverURL := range serverURLs {
		endpoints = append(endpoints, testServerHost(t, serverURL))
	}

	transport, err := newActivePassiveTransport(ActivePassiveClientConfig{
		Endpoints:        endpoints,
		BasePath:         "/",
		Schemes:          []string{"http"},
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

func testServerHost(t *testing.T, serverURL string) string {
	t.Helper()

	parsed, err := url.Parse(serverURL)
	if err != nil {
		t.Fatalf("parse server URL %q: %v", serverURL, err)
	}

	return parsed.Host
}

func setTestCurrentActiveHost(transport *ActivePassiveTransport, host string) {
	transport.mu.Lock()
	defer transport.mu.Unlock()
	transport.currentActiveHost = host
}
