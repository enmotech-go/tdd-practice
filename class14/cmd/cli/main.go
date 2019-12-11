package main

import (
	"fmt"
	"log"
	"os"
	"github.com/enmotech-go/tdd-practice/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := poker.NewCLI(store, os.Stdin)

	game.PlayPoker()

}
