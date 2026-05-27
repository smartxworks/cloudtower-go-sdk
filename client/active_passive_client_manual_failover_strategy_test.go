package client

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestManualFailoverStrategyReturnsSignalWithoutRediscoveryOn307(t *testing.T) {
	var staleProbeCount int32
	var staleRequestCount int32
	stale := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			atomic.AddInt32(&staleProbeCount, 1)
			w.WriteHeader(http.StatusTemporaryRedirect)
		case "/test":
			atomic.AddInt32(&staleRequestCount, 1)
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

	transport := newTestActivePassiveTransportWithStrategy(t, FailoverStrategyManualFailover, stale.URL, active.URL)
	setTestCurrentActiveEndpoint(transport, stale.URL)

	_, err := transport.Submit(testActivePassiveOperation())
	if !errors.Is(err, ErrActivePassiveFailoverRequired) {
		t.Fatalf("Submit() error = %v, want ErrActivePassiveFailoverRequired", err)
	}
	if got := transport.CurrentActiveEndpoint(); got != "" {
		t.Fatalf("CurrentActiveEndpoint() = %q, want empty", got)
	}
	if got := atomic.LoadInt32(&staleProbeCount); got != 0 {
		t.Fatalf("stale probe count = %d, want 0", got)
	}
	if got := atomic.LoadInt32(&staleRequestCount); got != 1 {
		t.Fatalf("stale request count = %d, want 1", got)
	}
	if got := atomic.LoadInt32(&activeProbeCount); got != 0 {
		t.Fatalf("active probe count = %d, want 0", got)
	}
	if got := atomic.LoadInt32(&activeRequestCount); got != 0 {
		t.Fatalf("active request count = %d, want 0", got)
	}
}
