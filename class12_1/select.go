package class12_1

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {

	t1 := time.Now()

	http.Get(a)

	aSince := time.Since(t1)

	t2 := time.Now()

	http.Get(b)

	bSince := time.Since(t2)

	if aSince < bSince {
		return a
	}

	return b
}
