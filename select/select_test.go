package sel

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare server speeds", func(t *testing.T) {
		slowServer := createDelayedServer(20 * time.Millisecond)
		fastServer := createDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Errorf("expected no error, got %q", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("with timeout", func(t *testing.T) {
		server := createDelayedServer(11 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
