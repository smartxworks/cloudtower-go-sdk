package client

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestDefaultFailoverStrategyDiscoversAndCachesPreferredAddress(t *testing.T) {
	var passiveProbeCount int32
	var passiveRequestCount int32
	passive := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			atomic.AddInt32(&passiveProbeCount, 1)
			w.WriteHeader(http.StatusTemporaryRedirect)
		case "/test":
			atomic.AddInt32(&passiveRequestCount, 1)
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer passive.Close()

	var activeProbeCount int32
	var activeRequestCount int32
	active := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			atomic.AddInt32(&activeProbeCount, 1)
			w.WriteHeader(http.StatusOK)
		case "/test":
			atomic.AddInt32(&activeRequestCount, 1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("active"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer active.Close()

	transport := newTestActivePassiveTransport(t, passive.URL, active.URL)

	result, err := transport.Submit(testActivePassiveOperation())
	if err != nil {
		t.Fatalf("Submit() error = %v", err)
	}
	if result != "active" {
		t.Fatalf("Submit() result = %v, want active", result)
	}

	activeHost := testServerHost(t, active.URL)
	if got := transport.CurrentActiveHost(); got != activeHost {
		t.Fatalf("CurrentActiveHost() = %q, want %q", got, activeHost)
	}

	result, err = transport.Submit(testActivePassiveOperation())
	if err != nil {
		t.Fatalf("second Submit() error = %v", err)
	}
	if result != "active" {
		t.Fatalf("second Submit() result = %v, want active", result)
	}
	if got := atomic.LoadInt32(&passiveProbeCount); got != 1 {
		t.Fatalf("passive probe count = %d, want 1", got)
	}
	if got := atomic.LoadInt32(&activeProbeCount); got != 1 {
		t.Fatalf("active probe count = %d, want 1", got)
	}
	if got := atomic.LoadInt32(&passiveRequestCount); got != 0 {
		t.Fatalf("passive request count = %d, want 0", got)
	}
	if got := atomic.LoadInt32(&activeRequestCount); got != 2 {
		t.Fatalf("active request count = %d, want 2", got)
	}
}

func TestDefaultFailoverStrategyKeepsPreferredAddressOnRecognizableError(t *testing.T) {
	var probeCount int32
	var requestCount int32
	statusCode := int32(http.StatusBadRequest)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			atomic.AddInt32(&probeCount, 1)
			w.WriteHeader(http.StatusOK)
		case "/test":
			atomic.AddInt32(&requestCount, 1)
			w.WriteHeader(int(atomic.LoadInt32(&statusCode)))
			_, _ = w.Write([]byte("active"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	transport := newTestActivePassiveTransport(t, server.URL)

	_, err := transport.Submit(testActivePassiveOperation())
	if !errors.Is(err, errActivePassiveTestRecognizable) {
		t.Fatalf("Submit() error = %v, want recognizable error", err)
	}

	host := testServerHost(t, server.URL)
	if got := transport.CurrentActiveHost(); got != host {
		t.Fatalf("CurrentActiveHost() = %q, want %q", got, host)
	}

	atomic.StoreInt32(&statusCode, http.StatusOK)
	result, err := transport.Submit(testActivePassiveOperation())
	if err != nil {
		t.Fatalf("second Submit() error = %v", err)
	}
	if result != "active" {
		t.Fatalf("second Submit() result = %v, want active", result)
	}
	if got := atomic.LoadInt32(&probeCount); got != 1 {
		t.Fatalf("probe count = %d, want 1", got)
	}
	if got := atomic.LoadInt32(&requestCount); got != 2 {
		t.Fatalf("request count = %d, want 2", got)
	}
}

func TestDefaultFailoverStrategyReprobesAndRetriesOnceOn307(t *testing.T) {
	var staleRequestCount int32
	stale := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			w.WriteHeader(http.StatusTemporaryRedirect)
		case "/test":
			atomic.AddInt32(&staleRequestCount, 1)
			w.Header().Set("Location", "/ignored")
			w.WriteHeader(http.StatusTemporaryRedirect)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer stale.Close()

	var activeProbeCount int32
	var activeRequestCount int32
	active := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			atomic.AddInt32(&activeProbeCount, 1)
			w.WriteHeader(http.StatusOK)
		case "/test":
			atomic.AddInt32(&activeRequestCount, 1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("active"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer active.Close()

	transport := newTestActivePassiveTransport(t, stale.URL, active.URL)
	setTestCurrentActiveHost(transport, testServerHost(t, stale.URL))

	result, err := transport.Submit(testActivePassiveOperation())
	if err != nil {
		t.Fatalf("Submit() error = %v", err)
	}
	if result != "active" {
		t.Fatalf("Submit() result = %v, want active", result)
	}

	activeHost := testServerHost(t, active.URL)
	if got := transport.CurrentActiveHost(); got != activeHost {
		t.Fatalf("CurrentActiveHost() = %q, want %q", got, activeHost)
	}
	if got := atomic.LoadInt32(&staleRequestCount); got != 1 {
		t.Fatalf("stale request count = %d, want 1", got)
	}
	if got := atomic.LoadInt32(&activeProbeCount); got != 1 {
		t.Fatalf("active probe count = %d, want 1", got)
	}
	if got := atomic.LoadInt32(&activeRequestCount); got != 1 {
		t.Fatalf("active request count = %d, want 1", got)
	}
}

func TestDefaultFailoverStrategyClearsPreferredAddressOnUnrecognizableError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			w.WriteHeader(http.StatusOK)
		case "/test":
			w.WriteHeader(http.StatusTeapot)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	transport := newTestActivePassiveTransport(t, server.URL)

	_, err := transport.Submit(testActivePassiveOperation())
	if err == nil {
		t.Fatal("Submit() error = nil, want unrecognizable error")
	}
	if got := transport.CurrentActiveHost(); got != "" {
		t.Fatalf("CurrentActiveHost() = %q, want empty", got)
	}
}

func TestDefaultFailoverStrategyReturnsRetryExhaustedAfterSecond307(t *testing.T) {
	stale := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			w.WriteHeader(http.StatusTemporaryRedirect)
		case "/test":
			w.Header().Set("Location", "/ignored")
			w.WriteHeader(http.StatusTemporaryRedirect)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer stale.Close()

	active := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			w.WriteHeader(http.StatusOK)
		case "/test":
			w.Header().Set("Location", "/ignored")
			w.WriteHeader(http.StatusTemporaryRedirect)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer active.Close()

	transport := newTestActivePassiveTransport(t, stale.URL, active.URL)
	setTestCurrentActiveHost(transport, testServerHost(t, stale.URL))

	_, err := transport.Submit(testActivePassiveOperation())
	if !errors.Is(err, ErrActivePassiveRetryExhausted) {
		t.Fatalf("Submit() error = %v, want ErrActivePassiveRetryExhausted", err)
	}
	if got := transport.CurrentActiveHost(); got != "" {
		t.Fatalf("CurrentActiveHost() = %q, want empty", got)
	}
}
