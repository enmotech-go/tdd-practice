package SelectJop

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	//slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	time.Sleep(20 * time.Millisecond)
	//	w.WriteHeader(http.StatusOK)
	//}))
	//
	//fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//}))

	//slowServer:=makeHttpServer(9 * time.Second)
	//fastServer:=makeHttpServer(0)
	//slowUrl := slowServer.URL
	//fastUrl := fastServer.URL
	//want := fastUrl
	//got := Racer(slowUrl,fastUrl)
	//if got !=want{
	//	t.Errorf("got '%s,want '%s'",got,want)
	//}
	//slowServer.Close()
	//fastServer.Close()

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeHttpServer(25 * time.Millisecond)
		defer server.Close()
		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
	t.Run("fastest one", func(t *testing.T) {
		slowServer := makeHttpServer(2 * time.Second)
		fastServer := makeHttpServer(1 * time.Second)
		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)
		if got != want {
			t.Errorf("got '%s,want '%s'", got, want)
		}
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		slowServer.Close()
		fastServer.Close()

	})

}

func makeHttpServer(delay time.Duration) *httptest.Server {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return mockServer
}
