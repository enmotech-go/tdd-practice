package main

import (
	poker "baikai/tdd-practice/server"
	"fmt"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

//type ReadWriteSeekTruncate interface {
//	io.ReadWriteSeeker
//	Truncate(size int64) error
//}

type ReadSeeker interface {
	Reader
	Seeker
}
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
	//game := poker.NewCLI(store,os.Stdin)
	//game.PlayPoker()
}
