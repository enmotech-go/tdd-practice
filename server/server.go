package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}
type ReadWriteSeekTruncate interface {
	io.ReadWriteSeeker
	Truncate(size int64) error
}
type PlayerServer struct {
	store PlayerStore
	http.Handler
}
type FileSystemStore struct {
	database *json.Encoder
	league League
}

func NewFileSystemStore(file *os.File)  (*FileSystemStore, error){
	file.Seek(0, 0)

	info ,err := file.Stat()
	if err!=nil{
		return nil,fmt.Errorf("problem getting file info from file %s, %v",file.Name(),err)
	}
	if info.Size()==0{
		file.Write([]byte("[]"))
		file.Seek(0,0)
	}
	league, err := NewLeague(file)
	if err!=nil{
		return nil,fmt.Errorf("problem loading player store from file %s, %v",file.Name(),err)
	}
	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	},nil
}

func (f *FileSystemStore) GetLeague() League {

	return f.league
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

type Player struct {
	Name string
	Wins int
}


func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
	//w.WriteHeader(http.StatusOK)
}
func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Chris", 20},
	}
}
func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
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
//func NewInMemoryPlayerStore() *InMemoryPlayerStore {
//	return &InMemoryPlayerStore{map[string]int{}}
//}
//
//type InMemoryPlayerStore struct {
//	store map[string]int
//}
//
//func (i *InMemoryPlayerStore) GetLeague() []Player {
//	var league []Player
//	for name, wins := range i.store {
//		league = append(league, Player{name, wins})
//	}
//	return league
//}
//
//func (i *InMemoryPlayerStore) RecordWin(name string) {
//	i.store[name]++
//}
//
//func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
//	return i.store[name]
//}
