package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	server := &PlayerServer{
		store: &StubPlayerStore{
			map[string]int{
				"Pepper": 20,
				"Floyd":  10,
			},
		},
	}

	t.Run("test_return_pepper_score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assert.Equal(t, "20", response.Body.String())
	})
	t.Run("test_return_floyd_score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assert.Equal(t, "10", response.Body.String())
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}
