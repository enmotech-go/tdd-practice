package selectpractise

import (
	"net/http"
"time"
	)

func Racer(a, b string) string {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)
	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	startA:= time.Now()
	http.Get(url)
	return time.Since(startA)
}
