package _select

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) string {
	durationA := measureResponseTime(urlA)
	durationB := measureResponseTime(urlB)

	if durationA < durationB {
		return urlA
	} else {
		return urlB
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
