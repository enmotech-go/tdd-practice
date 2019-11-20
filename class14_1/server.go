package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Player struct {
	Name string
	Wins int
}
type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	return nil
}
func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	if err := http.ListenAndServe(":15000", server); err != nil {
		log.Fatal("could not listen on port 5000 ", err)
		return
	}
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandle))
	router.Handle("/player/", http.HandlerFunc(p.playerHandle))
	p.Handler = router
	return p
}

type StubPlayerStore struct {
	score    map[string]int
	winCalls []string
	League   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.score[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.League
}

//func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//
//	router := http.NewServeMux()
//	router.Handle("/league", http.HandlerFunc(p.leagueHandle))
//	router.Handle("/player/", http.HandlerFunc(p.playerHandle))
//	router.ServeHTTP(w, r)
//}

func (p *PlayerServer) leagueHandle(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)

}

func (p *PlayerServer) playerHandle(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/player/"):]
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(player, w)
	}
}

func (p *PlayerServer) showScore(player string, w http.ResponseWriter) {

	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)

}
