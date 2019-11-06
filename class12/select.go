package class12

import (
	"context"
	"errors"
	"net/http"
	"time"
)

const DefaultTimeout = 10 * time.Second

func Racer(ctx context.Context, a, b string) (string, error) {
	return ConfigurableRacer(ctx, a, b, DefaultTimeout)
}

func ConfigurableRacer(ctx context.Context, a, b string, delay time.Duration) (string, error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(delay):
		return "", errors.New("timeout")
	case <-ctx.Done():
		return "", errors.New("cancel")
	}

}

func ping(url string) chan struct{} {
	result := make(chan struct{})
	go func() {
		http.Get(url)
		result <- struct{}{}
	}()
	return result
}

func measureResponseTime(url string) (duration time.Duration) {
	ta := time.Now()
	defer func() {
		duration = time.Now().Sub(ta)
	}()
	http.Get(url)
	return
}
