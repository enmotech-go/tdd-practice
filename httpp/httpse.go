package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}


type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league League
}

func (s *StubPlayerStore) GetLeague() League {
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

type Player struct {
	Name string
	Wins int
}

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name==name {
			return &l[i]
		}
	}
	return nil
}

type FileSystemStore struct {
	database *json.Encoder
	league League
}

func NewFileSystemStore(file *os.File) *FileSystemStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)
	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:league,
	}
}

func (f *FileSystemStore) GetLeague() League {
	return f.league
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
}

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := NewFileSystemStore(db)
	server := NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("can not server on 5000 %v" ,err)
	}
}
