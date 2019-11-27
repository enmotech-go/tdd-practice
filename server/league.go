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
	Read(p []byte)(n int,err error)
}
type Seeker interface {
	Seek(offset int64,whence int)(int64,error)
}
func NewLeague(rdr io.Reader)([]Player,error)  {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

