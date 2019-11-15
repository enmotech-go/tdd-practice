package selector

import (
	"testing"
	"time"
	"net/http/httptest"
	"net/http"
)

// const (
// 	slowURL = "http://www.a.com"
// 	fastURL = "http://www.b.com"
// )

func TestRacer(t *testing.T)  {
	t.Run("compare tow url of response speed", func(t *testing.T){
		slowServer := makeDelayedServer(time.Millisecond*20)
		fastServer := makeDelayedServer(time.Millisecond*0)
		defer func ()  {
			slowServer.Close()
			fastServer.Close()
		}()
	
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		got,err := Racer(slowURL, fastURL)
		if err != nil {
			t.Error(" expected an error but get one")
		}
		assertEqual(t, fastURL, got)
	})
	
	t.Run("returns an err if a doesn't response within 10s", func(t *testing.T){
		server := makeDelayedServer(time.Millisecond*15)
		defer server.Close()
		
	
		_,err := ConfigurableRacer(server.URL, server.URL, 12*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}


func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}
