package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var threeSecondTimeout = 3 * time.Second

func TestRacer(t *testing.T) {
	t.Run("v0", func(t *testing.T) {

		fastServer := makeDelayedServer(0 * time.Microsecond)
		slowServer := makeDelayedServer(60 * time.Microsecond)

		defer fastServer.Close()
		defer slowServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("slow: %s fast: %s \nwant: %s  got: %s", slowURL, fastURL, want, got)
		}
	})

	t.Run("return err if timeout 3s", func(t *testing.T) {
		serverA := makeDelayedServer(5 * time.Second)
		serverB := makeDelayedServer(2 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, threeSecondTimeout)
		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
