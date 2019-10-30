package class11

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) string {
	startA := time.Now()
	http.Get(url1)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(url2)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return url1
	}

	return url2
}
