package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}
//func PlayerServer(w http.ResponseWriter, r *http.Request) {
//	player := r.URL.Path[len("/players/"):]
//
//	fmt.Fprint(w, GetPlayerScore(player))
//}
func NewPlayerServer(store PlayerStore) *PlayerServer  {
	p :=new(PlayerServer)
	p.store=store
	router:=http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler=router
	return p
}
//func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	p.ServeHTTP(w,r)
//}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
}
func (p *PlayerServer)playersHandler(w http.ResponseWriter,r *http.Request)  {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
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
//
//func GetPlayerScore(name string) string {
//	if name == "Pepper" {
//		return "20"
//	}
//
//	if name == "Floyd" {
//		return "10"
//	}
//
//	return ""
//}
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct{
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}


func main() {

	server := NewPlayerServer(NewInMemoryPlayerStore())

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}