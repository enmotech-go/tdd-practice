package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	t.Run("return Pepper Score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/player/Pepper", nil)
		recorder := httptest.NewRecorder()
		PlayerServer(recorder, request)
		got := recorder.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got  '%s'  want '%s' ", got, want)
		}
	})

}
