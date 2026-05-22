package client

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestAlwaysProbeStrategyProbesBeforeEveryRequest(t *testing.T) {
	var probeCount int32
	var requestCount int32
	active := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			atomic.AddInt32(&probeCount, 1)
			w.WriteHeader(http.StatusOK)
		case "/test":
			atomic.AddInt32(&requestCount, 1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("active"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer active.Close()

	transport := newTestActivePassiveTransportWithStrategy(t, FailoverStrategyAlwaysProbe, active.URL)

	for i := 0; i < 2; i++ {
		result, err := transport.Submit(testActivePassiveOperation())
		if err != nil {
			t.Fatalf("Submit(%d) error = %v", i, err)
		}
		if result != "active" {
			t.Fatalf("Submit(%d) result = %v, want active", i, result)
		}
	}

	endpoint := active.URL
	if got := transport.CurrentActiveEndpoint(); got != endpoint {
		t.Fatalf("CurrentActiveEndpoint() = %q, want %q", got, endpoint)
	}
	if got := atomic.LoadInt32(&probeCount); got != 2 {
		t.Fatalf("probe count = %d, want 2", got)
	}
	if got := atomic.LoadInt32(&requestCount); got != 2 {
		t.Fatalf("request count = %d, want 2", got)
	}
}

func TestAlwaysProbeStrategyReturnsSignalOn307AfterFreshProbe(t *testing.T) {
	var probeCount int32
	var requestCount int32
	active := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/healthz":
			atomic.AddInt32(&probeCount, 1)
			w.WriteHeader(http.StatusOK)
		case "/test":
			atomic.AddInt32(&requestCount, 1)
			w.WriteHeader(http.StatusTemporaryRedirect)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer active.Close()

	transport := newTestActivePassiveTransportWithStrategy(t, FailoverStrategyAlwaysProbe, active.URL)

	_, err := transport.Submit(testActivePassiveOperation())
	if !errors.Is(err, ErrActivePassiveFailoverRequired) {
		t.Fatalf("Submit() error = %v, want ErrActivePassiveFailoverRequired", err)
	}
	if got := transport.CurrentActiveEndpoint(); got != "" {
		t.Fatalf("CurrentActiveEndpoint() = %q, want empty", got)
	}
	if got := atomic.LoadInt32(&probeCount); got != 1 {
		t.Fatalf("probe count = %d, want 1", got)
	}
	if got := atomic.LoadInt32(&requestCount); got != 1 {
		t.Fatalf("request count = %d, want 1", got)
	}
}
