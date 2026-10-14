package internal

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/phrase/phrase-go/v4"
)

func TestDownloadRetries(t *testing.T) {
	downloads := map[string]func(url, path string) error{
		"locale download": func(url, path string) error {
			cfg := phrase.NewConfiguration()
			cfg.BasePath = url
			file, _, err := downloadLocale(phrase.NewAPIClient(cfg), "proj1", "en", &phrase.LocaleDownloadOpts{})
			if err != nil {
				return err
			}
			return copyToDestination(file, path)
		},
		"async export download": downloadExportedLocale,
	}
	cases := []struct {
		name         string
		aborts       int32 // requests dropped mid-body before responding with status
		status       int
		wantErr      bool
		wantRequests int32
	}{
		{"retries interrupted body", 2, http.StatusOK, false, 3},
		{"gives up after max attempts", 3, http.StatusOK, true, 3},
		{"does not retry API errors", 0, http.StatusServiceUnavailable, true, 1},
	}

	prevAuth, prevConfig, prevSleep := Auth, Config, downloadRetrySleep
	Auth, Config, downloadRetrySleep = context.Background(), &phrase.Config{}, func(time.Duration) {}
	t.Cleanup(func() { Auth, Config, downloadRetrySleep = prevAuth, prevConfig, prevSleep })

	for name, download := range downloads {
		for _, c := range cases {
			t.Run(name+"/"+c.name, func(t *testing.T) {
				var requests int32
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if atomic.AddInt32(&requests, 1) <= c.aborts {
						abortBody(t, w)
						return
					}
					w.WriteHeader(c.status)
					io.WriteString(w, "new")
				}))
				defer srv.Close()
				path := filepath.Join(t.TempDir(), "en.json")
				os.WriteFile(path, []byte("old"), 0o644)

				err := download(srv.URL, path)

				if (err != nil) != c.wantErr {
					t.Errorf("expected error: %v, got %v", c.wantErr, err)
				}
				if requests != c.wantRequests {
					t.Errorf("expected %d requests, got %d", c.wantRequests, requests)
				}
				want := "new"
				if c.wantErr {
					want = "old"
				}
				if data, _ := os.ReadFile(path); string(data) != want {
					t.Errorf("expected file content %q, got %q", want, data)
				}
			})
		}
	}
}

// abortBody announces a body but drops the connection before sending it all,
// so the client fails while reading the response body.
func abortBody(t *testing.T, w http.ResponseWriter) {
	w.Header().Set("Content-Length", "100")
	io.WriteString(w, "partial")
	w.(http.Flusher).Flush()
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		t.Errorf("hijack failed: %v", err)
		return
	}
	conn.Close()
}
