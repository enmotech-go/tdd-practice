package main

import (
	"fmt"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":15000", server); err != nil {
		log.Fatal("could not listen on port 5000 ", err)
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

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.processWin(r, w)
	case http.MethodGet:
		p.showScore(r, p, w)
	}

}

func (p *PlayerServer) showScore(r *http.Request, p2 *PlayerServer, w http.ResponseWriter) {
	player := r.URL.Path[len("/player/"):]
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)

	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(r *http.Request, w http.ResponseWriter) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)

	}
}
