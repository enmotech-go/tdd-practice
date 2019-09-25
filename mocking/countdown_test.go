package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
go!`
	if got != want {
		t.Errorf("want %s got %s", want, got)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("no enough calls sleep, want 4 but %d", spySleeper.Calls)
	}
}
