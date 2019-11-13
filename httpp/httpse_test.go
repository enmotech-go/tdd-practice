package main

import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAA(t *testing.T) {
	newRequset := func(name string) *http.Request {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
		return  request
	}

	t.Run("return people score", func(t *testing.T) {
		request := newRequset("Pepper")
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"
		assert.Equal(t, got, want)
	})

	t.Run("return floyd's score", func(t *testing.T) {
		request := newRequset("floyd")
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "10"
		assert.Equal(t, got, want)
	})
}


