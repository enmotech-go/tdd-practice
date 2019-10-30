package selectpractise

import (
	"net/http"
	"errors"
	"time"
)

var senTimeout = 10 * time.Second

func Racer(a, b string) (string, error){
	return ConfableRacer(a, b, senTimeout)
}


func ConfableRacer(a, b string, duation time.Duration) (string, error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(duation):
		return "", errors.New("more than 10 ms")
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

