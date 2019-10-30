package websiteracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("quick test", func(t *testing.T) {
		ServerA := makeDelayedServer(20 * time.Millisecond)
		ServerB := makeDelayedServer(0 * time.Millisecond)

		defer ServerA.Close()
		defer ServerB.Close()

		A := ServerA.URL
		B := ServerB.URL

		want := B
		got, _ := Racer(A, B)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("slow test", func(t *testing.T) {
		ServerA := makeDelayedServer(30 * time.Millisecond)

		defer ServerA.Close()

		_, err := ConfigurableRacer(ServerA.URL, ServerA.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}
