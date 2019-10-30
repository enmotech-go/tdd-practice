package selectpractise

import (
	"time"
	"net/http"
	"net/http/httptest"
	"testing"
)

func makeDelayedServer(delay time.Duration) *httptest.Server  {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()
	slowUrl := slowServer.URL
	fastUrl := fastServer.URL
	want := fastUrl
	got := Racer(slowUrl, fastUrl)
	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}
