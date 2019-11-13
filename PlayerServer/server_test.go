package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {

	t.Run("return s Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "10"
		if got != want {
			t.Errorf("got '%s' , want '%s'", got, want)
		}
	})

}
