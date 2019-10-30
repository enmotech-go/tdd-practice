package select_racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
	t.Run("returns_error_if_timeout", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Second)
		fastServer := makeDelayedServer(10 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL)
		assert.Error(t, err)
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
