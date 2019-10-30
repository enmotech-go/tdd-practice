package class12_1

import (
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRear(t *testing.T) {
	slow := makeMockHttpServer(20 * time.Millisecond)
	fast := makeMockHttpServer(0 * time.Millisecond)
	defer slow.Close()
	defer fast.Close()
	slowURL := slow.URL
	fastURL := fast.URL
	want := fastURL
	got, err := Racer(slowURL, fastURL)
	if err != nil {
		t.Errorf("time out")
	}
	assert.Equal(t, got, want)
}

func makeMockHttpServer(duration time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
