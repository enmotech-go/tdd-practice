package main

import (
	"encoding/json"
	"fmt"
	"io"
)

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

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	var wins int
	for _,player := range f.GetLeague(){
		if player.Name==name{
			wins = player.Wins
			break
		}
	}
	return wins
}

func(f *FileSystemStore)RecordWin (name string)  {
	league :=f.GetLeague()
	for i , player :=range league{
		if player.Name==name{
			league[i].Wins++
		}
	}
	f.database.Seek(0,0)
	json.NewEncoder(f.database).Encode(league)
}