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
		server.ServeHTTP(recorder, request)
		assertResponseBody(recorder.Body.String(), "20", t)
		newGetRequestStart(t, recorder.Code, http.StatusOK)
	})

	t.Run("return Floyd Score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, request)
		assertResponseBody(recorder.Body.String(), "10", t)
		newGetRequestStart(t, recorder.Code, http.StatusOK)

	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		newGetRequestStart(t, response.Code, http.StatusNotFound)
	})
}

func newGetRequestStart(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}

}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return request
}

func newPostScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", name), nil)
	return request
}

func assertResponseBody(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		newGetRequestStart(t, response.Code, http.StatusAccepted)
	})

}

func TestName(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}
	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostScoreRequest(player)
		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, request)
		newGetRequestStart(t, recorder.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	newGetRequestStart(t, response.Code, http.StatusOK)
	assertResponseBody("3", response.Body.String(), t)
}

//func TestLeague(t *testing.T) {
//	store := StubPlayerStore{
//	}
//	server := &PlayerServer{&store}
//	t.Run("it retruns 200  /league", func(t *testing.T) {
//		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		newGetRequestStart(t, response.Code, http.StatusOK)
//	})
//
//}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := &PlayerServer{&store}

	t.Run("it retruns 200 on /league", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "?league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		newGetRequestStart(t, response.Code, http.StatusOK)

	})

}
