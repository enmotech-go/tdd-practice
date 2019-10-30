package class11

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) string {
	url1Duration := measureResponseTime(url1)
	url2Duration := measureResponseTime(url2)

	if url1Duration < url2Duration {
		return url1
	}

	return url2
}

func measureResponseTime(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}
