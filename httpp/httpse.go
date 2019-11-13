package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}
	if player == "floyd" {
		return "10"
	}
	return ""
}

func PlayerServer(recorder http.ResponseWriter, request *http.Request) {
	player := request.URL.Path[len("/player/"):]
	fmt.Fprintf(recorder, GetPlayerScore(player))
}

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("can not server on 5000 %v" ,err)
	}
}
