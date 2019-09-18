package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "chirs")
	got := buffer.String()
	want := "hello,chirs"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
