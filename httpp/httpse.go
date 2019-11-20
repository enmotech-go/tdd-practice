package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}


type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league []Player
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
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



func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

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

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
		i.store[name]++
		log.Println(i.store)
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	log.Println(name,i.store)
	return i.store[name]
}

type Player struct {
	Name string
	Wins int
}

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("can not server on 5000 %v" ,err)
	}
}
