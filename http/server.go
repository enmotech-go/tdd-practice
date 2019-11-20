package main

import "net/http"

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

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	if err := http.ListenAndServe(":5000", http.HandlerFunc(server.ServeHTTP)); err != nil {
		return
	}
}
