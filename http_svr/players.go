package main
import (
	"net/http"
	"fmt"
)

// class 13

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w,player)
	}

	// if r.Method == http.MethodPost {
	// 	w.WriteHeader(http.StatusAccepted)
	// 	return
	// }

	// player := r.URL.Path[len("/players/"):]
	// // w.WriteHeader(404)
	// score := p.store.GetPlayerScore(player)
	// if score == 0 {
	// 	w.WriteHeader(404)
	// }
	// fmt.Fprint(w, score)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	// player := r.URL.Path[len("/players/"):]
	// w.WriteHeader(404)
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(404)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	// p.store.RecordWin("Bob")
	// player := r.URL.Path[len("/players/"):]
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)

}
// func GetPlayerScore(name string) (score string) {
//     if name == "Pepper" {
// 		score = "20"
//         return
//     }

//     if name == "Floyd" {
// 		score = "10"
//         return
//     }

//     return
// }

