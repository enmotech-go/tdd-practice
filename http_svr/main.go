package main
import (
	"fmt"
	"net/http"
	"log"
)

func main()  {
	http_svr()
}

type InMemoryPlayerStore struct{
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int  {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
    return &InMemoryPlayerStore{map[string]int{}}
}



func http_svr()  {
	fmt.Println("hi http_svr")
	server := &PlayerServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", server);err != nil {
		log.Fatalf("listen on 5000 is failed, %v", err)
	}
}