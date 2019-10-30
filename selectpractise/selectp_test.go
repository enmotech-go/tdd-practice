package selectpractise

import (
	"time"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	slowUrl := slowServer.URL
	fastUrl := fastServer.URL
	want := fastUrl
	got := Racer(slowUrl, fastUrl)
	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}
