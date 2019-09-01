package hello

import (
	"testing"
)

func TestGreet(t *testing.T) {
	got := Hello()
	want := "hello world"
	if got != want {
		t.Errorf("got '%s',but want '%s'", got, want)
	}
}
