package SelectJop

import (
	"fmt"
	"net/http"
	"time"
)

var timeSecond = 10*time.Second
func Racer(a, b string) (winner string,err error) {
	return ConfigurableRacer(a,b,timeSecond)

}
func ConfigurableRacer(a,b string,timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

//func measureResponseTime(url string) time.Duration {
//	start := time.Now()
//	http.Get(url)
//	return time.Since(start)
//}
func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
