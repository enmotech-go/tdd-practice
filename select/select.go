package _select

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (string, error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(time.Second * 10):
		return "", fmt.Errorf("response time exceed 10s")
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
