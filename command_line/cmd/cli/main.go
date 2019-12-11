package main

import (
	"fmt"
	"os"
	"log"
	poker "enmotech-go/tdd-practice/command_line"
)

const dbFileName = "../game.db.json"

func main()  {
	store, fClose, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	fClose()
	
	fmt.Println("Let's play poker!")
	fmt.Println("Type {Name} wins to record a win")

	// db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatalf("problem opening %s %v", dbFileName, err)
	// }
	// store, err = poker.NewFileSystemStore(db)
	// if err != nil {
	// 	log.Fatalf("problem creating file system player store, %v", err)
	// }

	// game := poker.CLI{store, os.Stdin,}
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}