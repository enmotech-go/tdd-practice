package main

import (
	"log"
	"net/http"
	"os"
	"tdd-practice/httpp"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemStore(db)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("can not server on 5000 %v", err)
	}
}
