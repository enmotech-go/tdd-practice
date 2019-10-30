package SelectJop

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T)  {
	//slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	time.Sleep(20 * time.Millisecond)
	//	w.WriteHeader(http.StatusOK)
	//}))
	//
	//fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//}))
	slowServer:=makeHttpServer(20 * time.Millisecond)
	fastServer:=makeHttpServer(0)
	slowUrl := slowServer.URL
	fastUrl := fastServer.URL
	want := fastUrl
	got := Racer(slowUrl,fastUrl)
	if got !=want{
		t.Errorf("got '%s,want '%s'",got,want)
	}
	slowServer.Close()
	fastServer.Close()
}

func makeHttpServer(delay time.Duration) *httptest.Server{
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return mockServer
}