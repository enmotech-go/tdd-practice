package PlayerServer

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	response := httptest.NewRecorder()
	PlayerServer(response, request)
	got := response.Body.String()
	want := "20"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
