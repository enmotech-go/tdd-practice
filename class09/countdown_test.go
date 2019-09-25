package main

import (
	"bytes"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{1 * time.Second}
	Countdown(buffer, sleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
