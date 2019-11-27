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
type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
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
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins

	}
	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	}else{
			league = append(league,Player{name,1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
