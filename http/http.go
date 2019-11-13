package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func PlayerServer(w *httptest.ResponseRecorder, r *http.Request) {

	fmt.Fprint(w, "20")
}