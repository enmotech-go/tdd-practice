package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
}

type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	scores map[string] int
}


func (s StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/player/"):]
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}


type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("can not server on 5000 %v" ,err)
	}
}
