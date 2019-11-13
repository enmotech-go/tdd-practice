package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAA(t *testing.T) {
	t.Run("return people score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/player/Pepper", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("return floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/player/floyd", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "10"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}


