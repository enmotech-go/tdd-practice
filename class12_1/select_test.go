package class12_1

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRear(t *testing.T) {
	slow := makeMockHttpServer(11 * time.Second)
	fast := makeMockHttpServer(12 * time.Second)
	defer slow.Close()
	defer fast.Close()
	slowURL := slow.URL
	fastURL := fast.URL
	_, err := Racer(slowURL, fastURL)
	if err != nil {
		t.Log("time out")
	}
}

func makeMockHttpServer(duration time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
