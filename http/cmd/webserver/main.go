package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	http2 "tdd-practice/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []http2.Player {
	var league []http2.Player
	for name, wins := range i.store {
		league = append(league, http2.Player{name, wins})
	}
	return league
}

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := http2.NewFileSystemStore(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	server := http2.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", http.HandlerFunc(server.ServeHTTP)); err != nil {
		return
	}
}
