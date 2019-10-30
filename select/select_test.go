package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("response with in 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(time.Millisecond * 20)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0)
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("response exceed 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		defer serverA.Close()

		serverB := makeDelayedServer(12 * time.Second)
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)
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
