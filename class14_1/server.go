package main

import (
	"fmt"
	"net/http"
)

func main() {
	var server = &PlayerServer{store: &StubPlayerStore{score: map[string]int{}}}
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
