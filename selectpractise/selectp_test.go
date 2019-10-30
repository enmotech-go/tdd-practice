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
	t.Run("normal test ", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(12 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)
		if err != nil {
			t.Errorf("Reace err:%+v", err)
		}
		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}	
	})
	
	t.Run("return error great than 10 ms", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(12 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
		_, err := ConfableRacer(slowUrl, fastUrl, 100* time.Millisecond)
		if err != nil {
			t.Errorf("Reace err:%+v", err)
		}
	})
}
