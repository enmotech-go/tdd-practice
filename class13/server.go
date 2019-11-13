package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := &PlayerServer{&StubPlayerStore{scores: map[string]int{}}}
	if err := http.ListenAndServe(":5000", handler); err != nil {
		return
	}
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
