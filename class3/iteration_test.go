package class3

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"
	if got != want {
		t.Errorf("got  '%s' want '%s' ", got, want)
	}
}
