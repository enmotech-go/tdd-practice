package class02

import (
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	want := 4

	if sum != want {
		t.Errorf("want '%d' but got '%d'", want, sum)
	}
}
