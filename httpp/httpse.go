package httpp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func PlayerServer(recorder *httptest.ResponseRecorder, request *http.Request) {
	fmt.Fprintf(recorder, "20")
}
