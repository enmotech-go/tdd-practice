package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"
	if got != want {
		t.Errorf("want %s ,but get %s", want, got)
	}

}
