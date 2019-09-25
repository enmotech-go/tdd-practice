package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpyCountdownOP struct {
	Calls []string
}

func (s *SpyCountdownOP) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOP) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOP{})

		got := buffer.String()
		want := `3
2
1
go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("sleep after print", func(t *testing.T) {
		spySleeper := &SpyCountdownOP{}
		Countdown(spySleeper, spySleeper)
		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}
		if !reflect.DeepEqual(spySleeper.Calls, want) {
			t.Errorf("want %s got %s", want, spySleeper.Calls)
		}
	})

}
