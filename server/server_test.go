package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil, nil,
	}
	server := NewPlayerServer(&store)

	t.Run("test_return_pepper_score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "20", response.Body.String())
	})
	t.Run("test_return_floyd_score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "10", response.Body.String())
	})
	t.Run("test_404_on_missing_players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{map[string]int{}, []string{}, []Player{}}
	server := NewPlayerServer(&store)

	t.Run("test_records_wins_when_post", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusAccepted, response.Code)
		assert.Len(t, store.winCalls, 1)
		assert.Equal(t, player, store.winCalls[0])
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("test_get_score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "3", response.Body.String())
	})
	t.Run("test_get_league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assert.Equal(t, http.StatusOK, response.Code)
		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assert.Equal(t, want, got)
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func TestLeague(t *testing.T) {
	t.Run("test_return_200_on_/league", func(t *testing.T) {
		store := StubPlayerStore{}
		server := NewPlayerServer(&store)

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("test_returns_the_league_table_as_JSON", func(t *testing.T) {
		want := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, want}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("content-type"))
		assert.Equal(t, want, got)
	})
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	assert.NoError(t, err)
	return
}

func TestFileSystemStore(t *testing.T) {
	initialData := `
	[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}
	]
	`

	t.Run("league_from_a_reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store := FileSystemStore{database}
		got := store.GetLeague()
		want := League{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assert.Equal(t, want, got)

		got = store.GetLeague()
		assert.Equal(t, want, got)
	})
	t.Run("get_player_score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store := FileSystemStore{database}
		got := store.GetPlayerScore("Chris")
		want := 33
		assert.Equal(t, want, got)
	})
	t.Run("store_wins_for_existing_players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store := FileSystemStore{database}
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		assert.Equal(t, want, got)
	})
}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}
