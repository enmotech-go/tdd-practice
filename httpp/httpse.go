package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	scores map[string] int
	winCalls []string
}


func (s StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (p *PlayerServer) ProcessWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) ShowScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	log.Println(score)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}


func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/player/"):]
	switch r.Method {
	case http.MethodPost:
		p.ProcessWin(w, player)
	case http.MethodGet:
		p.ShowScore(w, player)
	}
}


func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct{
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
		i.store[name]++
		log.Println(i.store)
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	log.Println(name,i.store)
	return i.store[name]
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("can not server on 5000 %v" ,err)
	}
}
