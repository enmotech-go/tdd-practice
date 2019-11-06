package select 

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeMockHttpServer(cxt context.Context, delay time.Duration) *httptest.Server {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(delay):
		case <-cxt.Done():
		}
		// time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return mockServer
}

func TestRace(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx, cancelFun := context.WithCancel(context.Background())
		defer cancelFun()
		slowServer := makeMockHttpServer(ctx, 20*time.Millisecond)
		fastServer := makeMockHttpServer(ctx, 11*time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(ctx, slowServer.URL, fastServer.URL)
		assert.Equal(t, want, got)
		assert.NoError(t, err)

	})

	t.Run("timeout", func(t *testing.T) {
		ctx, cancelFun := context.WithCancel(context.Background())
		slowServer := makeMockHttpServer(ctx, 20*time.Second)
		defer slowServer.Close()

		_, err := ConfigurableRacer(ctx, slowServer.URL, slowServer.URL, 10*time.Millisecond)
		assert.NotNil(t, err)
		cancelFun()

	})

}
