package main

import (
	"fmt"
	"log"
	"os"

	poker "enmotech-go/tdd-practice/class13"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
