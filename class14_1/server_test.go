package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestServer(t *testing.T) {
	stub := StubPlayerStore{score: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}
	server := NewPlayerServer(&stub)
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
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/player/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		newGetRequestStart(t, response.Code, http.StatusAccepted)
	})

}

func TestName(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)
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

	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()
	store := NewFileSystemStore(database)
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	t.Run("get score ", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		newGetRequestStart(t, response.Code, http.StatusOK)
		assertResponseBody("3", response.Body.String(), t)
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		newGetRequestStart(t, response.Code, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
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

const jsonContentType = "application/json"

func TestLeague(t *testing.T) {
	t.Run("it retruns 200 on /league", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := StubPlayerStore{League: wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertContentType(response, t, jsonContentType)

		var got []Player
		got = getLeagueFromResponse(t, response.Body)
		newGetRequestStart(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})

}

func assertContentType(response *httptest.ResponseRecorder, t *testing.T, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.HeaderMap)
	}
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	if err := json.NewDecoder(body).Decode(&league); err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}
	return
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v' ", got, want)
	}

}

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store := NewFileSystemStore(database)
		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store := NewFileSystemStore(database)
		got := store.GetPlayerScore("Chris")

		want := 33

		assertScoreEquals(got, want, t)
	})

	t.Run("store wins for existing players", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store := NewFileSystemStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(got, want, t)

	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store := NewFileSystemStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(got, want, t)
	})

}

func assertScoreEquals(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got '%d' want '%d' ", got, want)
	}
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile

}
