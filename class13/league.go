package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(reader io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(reader).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
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
	} else {
		league = append(league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
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
