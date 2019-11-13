package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", handler); err != nil {
		return
	}
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
