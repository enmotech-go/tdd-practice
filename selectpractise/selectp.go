package selectpractise

import (
	"net/http"
)

func Racer(a, b string) string {
	select {
	case <- ping(a):
		return a
	case <- ping(b):
		return b
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

