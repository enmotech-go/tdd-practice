package main
import (
	"fmt"
	"net/http"
	"log"
)

func main()  {
	http_svr()
}

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int  {
	return 123
}

func http_svr()  {
	fmt.Println("hi http_svr")
	server := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server);err != nil {
		log.Fatalf("listen on 5000 is failed, %v", err)
	}
}