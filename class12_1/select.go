package class12_1

import (
	"fmt"
	"net/http"
	"time"
)

const DefaultTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {

	return ConfigurableRacer(a, b, DefaultTimeout)
}

func ConfigurableRacer(a, b string, delay time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(delay):
		return "", fmt.Errorf("time out")
	}

}

func measureResponseTime(url string) time.Duration {
	t1 := time.Now()
	http.Get(url)
	aSince := time.Since(t1)
	return aSince
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
