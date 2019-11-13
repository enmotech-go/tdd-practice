package main

import (
	"fmt"
	"log"
	"net/http"
)

func PlayerServer(recorder http.ResponseWriter, request *http.Request) {
	player := request.URL.Path[len("/player/"):]
	if player == "Pepper" {
		fmt.Fprintf(recorder, "20")
	}
	if player == "floyd" {
		fmt.Fprintf(recorder, "10")
	}
}

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("can not server on 5000 %v" ,err)
	}
}
