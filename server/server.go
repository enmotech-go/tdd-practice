package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
