package class12

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(30 * time.Second):
		return "", fmt.Errorf("time out")
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true

	}()
	return ch
}
