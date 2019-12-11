package main

import (
	"fmt"
	"log"
	"os"
	poker "tdd-practice/httpp"
)

const dbFileName = "../game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
