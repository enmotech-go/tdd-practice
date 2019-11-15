package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		// spySleeper := &SpySleeper{}
		Countdown(buffer, &CountdownOperationsSpy{})
		want := `3
2
1
Go!`
		assertEqual(t, buffer.String(), want)
		// if spySleeper.Calls != 4 {
		// 	t.Errorf("not enough calls to sleeper, got %d, want 4", spySleeper.Calls)
		// }

	})
	t.Run("write after sleep!", func(t *testing.T) {
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
			t.Errorf("want calss %v, got %v", want, spySleepPrinter.Calls)
		}
	})

}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
