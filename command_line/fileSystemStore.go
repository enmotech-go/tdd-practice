package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemStore struct {
	// database io.Reader
	// database io.ReadSeeker
	// database io.ReadWriteSeeker
	database *json.Encoder
	league   League
}

func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	// file.Seek(0, 0)
	// info, err := file.Stat()

	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	// if info.Size() == 0 {
	// 	file.Write([]byte("[]"))
	// 	file.Seek(0, 0)
	// }

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemStore{
		// database: &tape{database},
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func (f *FileSystemStore) GetLeague() League {
	// f.database.Seek(0, 0)
	// league, _ := NewLeague(f.database)
	// league := f.league

	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	// var wins int
	// for _, player := range f.GetLeague() {
	// 	if player.Name == name {
	// 		wins = player.Wins
	// 		break
	// 	}
	// }
	// player := f.GetLeague().Find(name)
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	// league := f.GetLeague()

	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	// for i, player := range league {
	// 	if player.Name == name {
	// 		league[i].Wins ++
	// 	}
	// }

	// f.database.Seek(0, 0)
	// json.NewEncoder(f.database).Encode(league)
	f.database.Encode(f.league)

}

type League []Player

func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
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

type tape struct {
	// file io.ReadWriteSeeker
	file *os.File
}

// 当写入时，从头部开始
func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}

type ReadWriteSeekTruncate interface {
	io.ReadWriteSeeker
	Truncate(size int64) error
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}
