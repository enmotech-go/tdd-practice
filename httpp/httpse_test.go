package main

import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRequset(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return  request
}

func TestAA(t *testing.T) {
	server := &PlayerServer{&StubPlayerStore{map[string]int{
		"Pepper": 20,
		"floyd":  10,
	}, nil}}

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

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("it records wins when POST", func(t *testing.T) {
		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t,http.StatusAccepted, response.Code)
		assert.Equal(t, 1, len(store.winCalls))
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", name), nil)
	return req
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newRequset(player))
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, "3", response.Body.String())
}
