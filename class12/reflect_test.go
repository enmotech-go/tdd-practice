package class12

import (
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("got '%d' want '%d'", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got '%s' want '%s'", got[0], expected)
	}
}
