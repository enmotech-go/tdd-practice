package selector

import (
	"time"
	"net/http"
	"fmt"
)

/**
class_11
demand: 比较两个网站的相应速度，快的返回，如果10内未相应则返回error
*/

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error)  {
	// aDuration := measureResponseTime(a)
	// bDuration := measureResponseTime(b)
	// if aDuration < bDuration {
	// 	return a
	// }
	// return b
	
	return ConfigurableRacer(a,b,tenSecondTimeout)

}

func ConfigurableRacer(a,b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out for wait %s and %s", a, b)
	}
}

func measureResponseTime(url string) time.Duration  {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
