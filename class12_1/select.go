package class12_1

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {

	aSince := measureResponseTime(a)

	bSince := measureResponseTime(b)

	if aSince < bSince {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	t1 := time.Now()
	http.Get(url)
	aSince := time.Since(t1)
	return aSince
}
