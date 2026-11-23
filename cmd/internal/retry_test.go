package internal

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type rejection struct{ reason, reset string }

type rateLimitRun struct {
	status int
	calls  int
	slept  []time.Duration
}

// runRateLimited posts to a server that rejects request n with rejects(n) (nil
// accepts it), the clock standing at unix time 1000. Sleeps are recorded, not slept.
func runRateLimited(t *testing.T, rejects func(call int) *rejection) rateLimitRun {
	prevSleep, prevNow := retrySleep, rateLimitNow
	t.Cleanup(func() { retrySleep, rateLimitNow = prevSleep, prevNow })
	rateLimitNow = func() time.Time { return time.Unix(1000, 0) }
	var run rateLimitRun
	retrySleep = func(_ context.Context, d time.Duration) error { run.slept = append(run.slept, d); return nil }

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		if string(body) != "payload" {
			t.Errorf("body = %q, want payload", body)
		}
		run.calls++
		if rej := rejects(run.calls); rej != nil {
			w.Header().Set("X-Rate-Limit-Reason", rej.reason)
			w.Header().Set("X-Rate-Limit-Remaining", "0")
			if rej.reset != "" {
				w.Header().Set("X-Rate-Limit-Reset", rej.reset)
			}
			w.WriteHeader(http.StatusTooManyRequests)
		}
	}))
	defer srv.Close()

	client := &http.Client{Transport: rateLimitTransport{http.DefaultTransport}}
	resp, err := client.Post(srv.URL, "text/plain", strings.NewReader("payload"))
	if err != nil {
		t.Fatal(err)
	}
	run.status = resp.StatusCode
	return run
}

func TestRateLimitTransport(t *testing.T) {
	conc := &rejection{"global-concurrency", ""}
	rate := &rejection{"global-rate", "1010"}
	// wanted sleep ranges: exact for waiting until reset, backoff plus up to 50% jitter otherwise
	exact := func(d time.Duration) [2]time.Duration { return [2]time.Duration{d, d} }
	backoff := func(d time.Duration) [2]time.Duration { return [2]time.Duration{d, d * 3 / 2} }

	cases := []struct {
		name    string
		rejects []*rejection
		want    [][2]time.Duration
	}{
		{"concurrency backs off exponentially", []*rejection{conc, conc, conc},
			[][2]time.Duration{backoff(250 * time.Millisecond), backoff(500 * time.Millisecond), backoff(time.Second)}},
		{"rate waits for reset", []*rejection{rate}, [][2]time.Duration{exact(11 * time.Second)}},
		{"reset in the past", []*rejection{{"global-rate", "900"}}, [][2]time.Duration{exact(concurrencyDelay)}},
		{"no reason uses reset", []*rejection{{"", "1005"}}, [][2]time.Duration{exact(6 * time.Second)}},
		{"no reason without reset backs off", []*rejection{{"", ""}}, [][2]time.Duration{backoff(concurrencyDelay)}},
		{"rate without reset backs off", []*rejection{{"global-rate", ""}}, [][2]time.Duration{backoff(concurrencyDelay)}},
		{"mixed reasons", []*rejection{conc, rate, conc},
			[][2]time.Duration{backoff(250 * time.Millisecond), exact(11 * time.Second), backoff(time.Second)}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			run := runRateLimited(t, func(call int) *rejection {
				if call <= len(c.rejects) {
					return c.rejects[call-1]
				}
				return nil
			})
			if run.status != 200 || run.calls != len(c.rejects)+1 {
				t.Errorf("status=%d calls=%d, want 200/%d", run.status, run.calls, len(c.rejects)+1)
			}
			if len(run.slept) != len(c.want) {
				t.Fatalf("slept %v, want %d sleeps", run.slept, len(c.want))
			}
			for i, d := range run.slept {
				if d < c.want[i][0] || d > c.want[i][1] {
					t.Errorf("sleep %d = %s, want within %v", i, d, c.want[i])
				}
			}
		})
	}
}

func TestRateLimitTransportGivesUp(t *testing.T) {
	run := runRateLimited(t, func(int) *rejection { return &rejection{"global-concurrency", ""} })
	if run.status != 429 || run.calls != maxRateLimitRetries+1 {
		t.Errorf("status=%d calls=%d, want 429/%d", run.status, run.calls, maxRateLimitRetries+1)
	}

	// a reset beyond the wait budget returns the 429 without waiting
	run = runRateLimited(t, func(int) *rejection { return &rejection{"global-rate", "4600"} })
	if run.status != 429 || run.calls != 1 || len(run.slept) != 0 {
		t.Errorf("status=%d calls=%d sleeps=%v, want 429 after 1 call without sleeping", run.status, run.calls, run.slept)
	}
}

func TestRateLimitTransportCancelledDuringBackoff(t *testing.T) {
	prevSleep := retrySleep
	t.Cleanup(func() { retrySleep = prevSleep })
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.Header().Set("X-Rate-Limit-Reason", "global-concurrency")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	retrySleep = func(ctx context.Context, d time.Duration) error {
		cancel()
		return sleepContext(ctx, time.Hour)
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, srv.URL, strings.NewReader("payload"))
	_, err := rateLimitTransport{http.DefaultTransport}.RoundTrip(req)
	if !errors.Is(err, context.Canceled) || calls != 1 {
		t.Errorf("err=%v calls=%d, want context.Canceled after 1 call", err, calls)
	}
}
