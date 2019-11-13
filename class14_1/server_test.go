package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	stub := StubPlayerStore{score: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}
	server := &PlayerServer{&stub}
	t.Run("return Pepper Score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		recorder := httptest.NewRecorder()
		server.ServerHTTP(recorder, request)
		assertResponseBody(recorder.Body.String(), "20", t)
		got := recorder.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got  '%s'  want '%s' ", got, want)
		}
	})
	t.Run("return Floyd Score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		recorder := httptest.NewRecorder()
		server.ServerHTTP(recorder, request)
		assertResponseBody(recorder.Body.String(), "10", t)

	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return request
}

func assertResponseBody(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
