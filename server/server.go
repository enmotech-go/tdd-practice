package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}
type FileSystemStore struct {
	database io.ReadSeeker
}

func (f *FileSystemStore)GetLeague()[]Player  {
	f.database.Seek(0,0)
	league ,_:=NewLeague(f.database)
	return league
}

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
type Player struct {
	Name string
	Wins int
}

//const  jsonContentType  =  "application/json"

func (p *PlayerServer) leagueHandler(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
	//w.WriteHeader(http.StatusOK)
}
func (p *PlayerServer) getLeagueTable()[]Player  {
	return []Player{
		{"Chris",20},
	}
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

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct{
	store map[string]int
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name ,wins :=range i.store{
		league=append(league,Player{name,wins})
	}
	return league
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

