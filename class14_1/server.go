package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Path[len("/player/"):]
	getPlayerScore(player, w)

}

func getPlayerScore(player string, w http.ResponseWriter) {
	if player == "Pepper" {
		fmt.Fprint(w, "20")
		return
	}
	if player == "Floyd" {
		fmt.Fprint(w, "10")
		return
	}
}

func main() {
	handlerFunc := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handlerFunc); err != nil {
		return
	}
}
