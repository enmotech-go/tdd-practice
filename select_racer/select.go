package select_racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureResposeTime(a)
	bDuration := measureResposeTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResposeTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
