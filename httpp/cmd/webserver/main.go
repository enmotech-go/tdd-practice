package main

import (
	"log"
	"net/http"
	"tdd-practice/httpp"
)

const dbFileName = "../game.db.json"

func main() {
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("can not server on 5000 %v", err)
	}
}
