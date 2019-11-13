package class13

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayersServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

func PlayersServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	if player == "Pepper" {
		fmt.Fprint(w, "20")
	}

	if player == "Floyd" {
		fmt.Fprint(w, "10")
	}
}
