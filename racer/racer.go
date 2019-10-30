package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measurerResponseTime(a)
	bDuration := measurerResponseTime(b)
	if aDuration < bDuration {
		return a
	}
	return b
}

func measurerResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
