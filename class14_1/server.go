package main

import (
	"fmt"
	"net/http"
)

func GetPlayerScore(player string) string {
	if player == "Pepper" {

		return "20"
	}
	if player == "Floyd" {

		return "10"
	}
	return ""
}

func main() {
	var server = &PlayerServer{}
	handlerFunc := http.HandlerFunc(server.ServerHTTP)
	if err := http.ListenAndServe(":5000", handlerFunc); err != nil {
		return
	}
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	score map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.score[name]
}

func (p *PlayerServer) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/player/"):]
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
