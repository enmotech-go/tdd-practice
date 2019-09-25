package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("sleep after every print", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}
		Countdown(buffer, sleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
		if sleeper.Calls != 4 {
			t.Errorf("want 4 got %d", sleeper.Calls)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

}
