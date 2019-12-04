package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}
type ReadWriteSeekTruncate interface {
	io.ReadWriteSeeker
	Truncate(size int64) error
}
type PlayerServer struct {
	store PlayerStore
	http.Handler
}
type FileSystemStore struct {
	database *json.Encoder
	league League
}

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
type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store ,err:= NewFileSystemStore(db)
	if err!=nil{
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
