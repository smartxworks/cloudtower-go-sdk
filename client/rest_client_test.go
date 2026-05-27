package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func TestHTTPClientWithConfigPreservesEndpointPathPrefix(t *testing.T) {
	var sawProbe bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/tower-a/api/healthz":
			sawProbe = true
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := NewHTTPClientWithConfig(strfmt.Default, &TransportConfig{
		CommonBasePath: server.URL + "/tower-a",
		BasePath:       "/v2/api",
	})

	transport, ok := client.Transport.(*httptransport.Runtime)
	if !ok {
		t.Fatalf("Transport = %T, want *httptransport.Runtime", client.Transport)
	}
	if want := "/tower-a/v2/api"; transport.BasePath != want {
		t.Fatalf("BasePath = %q, want %q", transport.BasePath, want)
	}

	active, err := client.ProbeActivePassive(context.Background())
	if err != nil {
		t.Fatalf("ProbeActivePassive() error = %v", err)
	}
	if !active {
		t.Fatal("ProbeActivePassive() = false, want true")
	}
	if !sawProbe {
		t.Fatal("probe did not use endpoint path prefix")
	}
}

func TestHTTPClientWithConfigUsesCommonBasePathWithHostAndSchemes(t *testing.T) {
	var sawProbe bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/tower-a/api/healthz":
			sawProbe = true
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	parsedURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse test server URL: %v", err)
	}

	client := NewHTTPClientWithConfig(strfmt.Default, &TransportConfig{
		Host:           parsedURL.Host,
		Schemes:        []string{parsedURL.Scheme},
		CommonBasePath: "/tower-a",
		BasePath:       "/v2/api",
	})

	transport, ok := client.Transport.(*httptransport.Runtime)
	if !ok {
		t.Fatalf("Transport = %T, want *httptransport.Runtime", client.Transport)
	}
	if want := "/tower-a/v2/api"; transport.BasePath != want {
		t.Fatalf("BasePath = %q, want %q", transport.BasePath, want)
	}

	active, err := client.ProbeActivePassive(context.Background())
	if err != nil {
		t.Fatalf("ProbeActivePassive() error = %v", err)
	}
	if !active {
		t.Fatal("ProbeActivePassive() = false, want true")
	}
	if !sawProbe {
		t.Fatal("probe did not use common base path")
	}
}

func TestHTTPClientWithConfigUsesCustomProbePath(t *testing.T) {
	var sawProbe bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/tower-a/custom/healthz":
			sawProbe = true
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := NewHTTPClientWithConfig(strfmt.Default, &TransportConfig{
		CommonBasePath: server.URL + "/tower-a",
		BasePath:       "/v2/api",
		ProbePath:      "/custom/healthz",
	})

	active, err := client.ProbeActivePassive(context.Background())
	if err != nil {
		t.Fatalf("ProbeActivePassive() error = %v", err)
	}
	if !active {
		t.Fatal("ProbeActivePassive() = false, want true")
	}
	if !sawProbe {
		t.Fatal("probe did not use custom probe path")
	}
}

func TestHTTPClientWithConfigKeepsLegacyHostBasePath(t *testing.T) {
	var sawProbe bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			sawProbe = true
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	parsedURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse test server URL: %v", err)
	}

	client := NewHTTPClientWithConfig(strfmt.Default, &TransportConfig{
		Host:     parsedURL.Host,
		BasePath: "/v2/api",
		Schemes:  []string{"http"},
	})

	transport, ok := client.Transport.(*httptransport.Runtime)
	if !ok {
		t.Fatalf("Transport = %T, want *httptransport.Runtime", client.Transport)
	}
	if transport.Host != parsedURL.Host {
		t.Fatalf("Host = %q, want %q", transport.Host, parsedURL.Host)
	}
	if transport.BasePath != "/v2/api" {
		t.Fatalf("BasePath = %q, want /v2/api", transport.BasePath)
	}

	active, err := client.ProbeActivePassive(context.Background())
	if err != nil {
		t.Fatalf("ProbeActivePassive() error = %v", err)
	}
	if !active {
		t.Fatal("ProbeActivePassive() = false, want true")
	}
	if !sawProbe {
		t.Fatal("legacy probe did not use default probe path")
	}
}
