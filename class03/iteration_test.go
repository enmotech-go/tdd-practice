package class03

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaa"

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}
