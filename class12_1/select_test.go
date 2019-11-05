package class12_1

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRear(t *testing.T) {

	t.Run("success ", func(t *testing.T) {
		slow := makeMockHttpServer(20 * time.Millisecond)
		fast := makeMockHttpServer(0)
		defer slow.Close()
		defer fast.Close()
		slowURL := slow.URL
		fastURL := fast.URL
		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Log("time out")
		}
		assert.Equal(t, got, want)
	})

	t.Run("no timeout ", func(t *testing.T) {
		slow := makeMockHttpServer(20 * time.Millisecond)
		fast := makeMockHttpServer(0)
		defer slow.Close()
		defer fast.Close()
		slowURL := slow.URL
		fastURL := fast.URL
		_, err := ConfigurableRacer(slowURL, fastURL, 1*time.Second)
		assert.NoError(t, err)
	})

	t.Run("timeout ", func(t *testing.T) {
		fast := makeMockHttpServer(1000 * time.Millisecond)
		fastURL := fast.URL
		_, err := ConfigurableRacer(fastURL, fastURL, 10*time.Millisecond)
		assert.Error(t, err)
		fast.Close()
	})

}

func makeMockHttpServer(duration time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
