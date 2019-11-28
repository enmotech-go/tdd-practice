package command_line

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const jsonContentType = "json/application"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	// GetLeague() []Player
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	// router *http.ServeMux
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	// p := &PlayerServer{
	// 	store,
	// 	http.NewServeMux(),
	// }
	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router
	return p
}

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// router := http.NewServeMux()

// 	// router.Handle("/league", http.HandlerFunc(p.leagueHandler))
// 	// router.Handle("/players/", http.HandlerFunc(p.playersHandler))
// 	// router.ServeHTTP(w, r)
// 	// if r.Method == http.MethodPost {
// 	// 	w.WriteHeader(http.StatusAccepted)
// 	// 	return
// 	// }
// 	// player := r.URL.Path[len("/players/"):]
// 	// // w.WriteHeader(404)
// 	// score := p.store.GetPlayerScore(player)
// 	// if score == 0 {
// 	// 	w.WriteHeader(404)
// 	// }
// 	// fmt.Fprint(w, score)

// 	p.router.ServeHTTP(w, r)
// }

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())

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
	// player := r.URL.Path[len("/players/"):]
	// w.WriteHeader(404)
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound) // did't exist player'score is 0
	}
	// score = 2
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



type Player struct {
	Name string
	Wins int
}
