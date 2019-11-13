package main

import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAA(t *testing.T) {
	server := &PlayerServer{&StubPlayerStore{map[string]int{
		"Pepper": 20,
		"floyd":  10,
	}}}

	newRequset := func(name string) *http.Request {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
		return  request
	}

	t.Run("return people score", func(t *testing.T) {
		request := newRequset("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "20"
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, want, got)
	})

	t.Run("return floyd's score", func(t *testing.T) {
		request := newRequset("floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "10"
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, want, got)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newRequset("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		assert.Equal(t, want, got)
	})
}


