package main

import "net/http"

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func Run() error {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", http.HandlerFunc(server.ServeHTTP)); err != nil {
		return err
	}
	return nil
}
