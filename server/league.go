package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins

	}
	return 0
}

func (f *FileSystemStore) RecordWin(name string) {

	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	}else{
			f.league = append(f.league,Player{name,1})
	}

	//f.database.Seek(0, 0)
	f.database.Encode(f.league)
}
