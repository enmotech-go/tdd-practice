package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	t.Run("return Pepper Score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		recorder := httptest.NewRecorder()
		PlayerServer(recorder, request)
		got := recorder.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got  '%s'  want '%s' ", got, want)
		}
	})
	t.Run("return Floyd Score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		recorder := httptest.NewRecorder()
		PlayerServer(recorder, request)
		got := recorder.Body.String()
		want := "10"
		if got != want {
			t.Errorf("got  '%s'  want '%s' ", got, want)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return request
}
