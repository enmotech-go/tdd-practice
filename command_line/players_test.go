package command_line

import (
	// "encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	// "strings"
	"testing"
	// "github.com/stretchr/testify/assert"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{"Pepper": 20, "Floyd": 10}, nil, nil,
	}
	// server := &PlayerServer{
	// 	&store,
	// 	nil,
	// }

	server := NewPlayerServer(&store)

	t.Run("returns Pepper`s score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Code
		// got = 404
		want := http.StatusNotFound

		assertStatus(t, got, want)
	})
}

// post
func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	// server := &PlayerServer{&store, nil}

	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		// request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		player := "Pepper" // todo
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got is %q, want is %q", store.winCalls[0], player)
		}

	})

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		// if len(store.winCalls) != 2 {
		// 	t.Fatalf("got %d calls to RecordWin, want %d", len(store.winCalls), 2)
		// }

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got '%s', want '%s'", store.winCalls[0], player)
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got is %s, want is %s", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did`t get correct status, got is %d, want is %d", got, want)
	}
}

const count = 4

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	// store := InMemoryPlayerStore{}
	// store := NewInMemoryPlayerStore()
	// server := PlayerServer{store, nil}
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()
	// store := &FileSystemStore{database}
	store, err := NewFileSystemStore(database)
	assertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	for i := 0; i < count; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)                      // 若没有请求则无存根
		assertResponseBody(t, response.Body.String(), strconv.Itoa(count)) // 3次Post请求，store["Pepper"]=3
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", count},
		}
		assertLeague(t, got, want)
	})
}

func TestLeague(t *testing.T) {

	// t.Run("it returns 200 on /league", func(t *testing.T) {
	// 	store := StubPlayerStore{}
	// 	server := &PlayerServer{&store, nil}

	// 	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	// 	response := httptest.NewRecorder()

	// 	server.ServeHTTP(response, request)

	// 	assertStatus(t, response.Code, http.StatusOK)
	// })

	t.Run("it return json league", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)

		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, jsonContentType)

	})
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()
	// err := json.NewDecoder(body).Decode(&league)
	league, err := NewLeague(body)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got is %v, want is %v", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func TestFileSystemStore(t *testing.T) {
	t.Run("league store", func(t *testing.T) {
		// database := strings.NewReader(`[
		//     {"Name": "Cleo", "Wins": 10},
		//     {"Name": "Chris", "Wins": 33}]`)
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		// store := FileSystemStore{database}
		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		// database := strings.NewReader(`[
		//     {"Name": "Cleo", "Wins": 10},
		//     {"Name": "Chris", "Wins": 33}]`)

		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		// store := FileSystemStore{database}
		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		// store := FileSystemStore{database}
		store, err := NewFileSystemStore(database)
		assertNoError(t, err)
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		// store := FileSystemStore{database}
		_, err := NewFileSystemStore(database)
		assertNoError(t, err)

	})

}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got is %d, want is %d", got, want)
	}
}

// func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
// 	t.Helper()
// 	tmpfile, err := ioutil.TempFile("", "db")
// 	if err != nil {
// 		t.Fatalf("could not create temp file %v", err)
// 	}

// 	tmpfile.Write([]byte(initialData))
// 	removeFile := func() {
// 		tmpfile.Close()
// 		os.Remove(tmpfile.Name())
// 	}

// 	return tmpfile, removeFile
// }
func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func TestTape_Write(t *testing.T) {
	t.Run("write........", func(t *testing.T) {
		file, clean := createTempFile(t, "12345")
		defer clean()

		tape := &tape{file}
		tape.Write([]byte("abc"))
		file.Seek(0, 0)

		newFileContents, _ := ioutil.ReadAll(file)
		got := string(newFileContents)
		want := "abc"
		if got != want {
			t.Errorf("got is '%s', want is '%s'", got, want)
		}
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect an error but got one, %v", err)
	}
}

// func typeAssert(file io.ReadWriteSeeker) *os.File  {

// }
