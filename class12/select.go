package main

import (
	"errors"
	"net/http"
	"time"
)

const DefaultTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, DefaultTimeout)
}

func ConfigurableRacer(a, b string, delay time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(delay):
		return "", errors.New("timeout")
	}

}

func ping(url string) chan struct{} {
	result := make(chan struct{}, 1)
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
