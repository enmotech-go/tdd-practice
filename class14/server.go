package poker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
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

//json
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(p.store.GetLeague()); err != nil {
		log.Println("leagueHandler Encode with err", err)
	}
	w.WriteHeader(http.StatusOK)
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

	score, ok := p.store.GetPlayerScore(player)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

type Player struct {
	Name string
	Wins int
}

//io
type FileSystemStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, err
	}
	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}
func initialisePlayerDBFile(file *os.File) error {
	_, err := file.Seek(0, 0)
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}
	if info.Size() == 0 {
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return err
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FileSystemStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemStore) GetPlayerScore(name string) (int, bool) {
	var wins int
	var exist bool
	player := f.league.Find(name)
	if player != nil {
		return player.Wins, true
	}
	return wins, exist
}

func (f *FileSystemStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	if err := f.database.Encode(f.league); err != nil {
		log.Println("RecordWin  encode with err", err)
	}
}
