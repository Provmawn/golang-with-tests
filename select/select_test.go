package sel

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(d time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}

func TestRacer(t *testing.T) {
	t.Run("returns faster server", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		got, _ := Racer(slowServer.URL, fastServer.URL, tenSecondTimeout)
		want := fastServer.URL

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns error if function is longer than 10 seconds", func(t *testing.T) {
		slowServer := makeDelayedServer(500 * time.Millisecond)
		fastServer := makeDelayedServer(500 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL, 500*time.Millisecond)
		if err == nil {
			t.Error("expected to get an error")
		}
	})
}
