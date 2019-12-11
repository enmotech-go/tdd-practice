package main

import (
	"fmt"
	"log"
	"os"
	poker "tdd-practice/server"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Let's play poker")
	fmt.Println("Type {name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
