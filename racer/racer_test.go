package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the b", func(t *testing.T) {
		aServer := makeDelayedServer(20 * time.Millisecond)
		bServer := makeDelayedServer(0 * time.Millisecond)

		defer aServer.Close()
		defer bServer.Close()

		aURL := aServer.URL
		bURL := bServer.URL

		want := bURL
		got, err := Racer(aURL, bURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("compares speeds of servers, returning the url of the a", func(t *testing.T) {
		aServer := makeDelayedServer(0 * time.Millisecond)
		bServer := makeDelayedServer(20 * time.Millisecond)

		defer aServer.Close()
		defer bServer.Close()

		slowURL := aServer.URL
		fastURL := bServer.URL

		want := slowURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
