package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeMockHttpServer(delay time.Duration) *httptest.Server {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return mockServer
}

func TestRace(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		slowServer := makeMockHttpServer(20 * time.Millisecond)

		fastServer := makeMockHttpServer(11 * time.Millisecond)

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)
		assert.Equal(t, want, got)
		assert.NoError(t, err)
		slowServer.Close()
		fastServer.Close()

	})

	t.Run("timeout", func(t *testing.T) {
		slowServer := makeMockHttpServer(20 * time.Second)
		fastServer := makeMockHttpServer(11 * time.Second)
		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 10*time.Millisecond)

		fastServer.Close()
		slowServer.Close()
		assert.NotNil(t, err)

	})

}
