package class12

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

func measureResponseTime(a string) time.Duration {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)
	return aDuration
}
