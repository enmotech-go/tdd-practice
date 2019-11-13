package class13

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayersServer(response, request)
		got := response.Body.String()
		want := "20"

		if got != want {
			t.FailNow()
		}
	})
}
