package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

func main() {
	handlerFunc := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handlerFunc); err != nil {
		return
	}
}
