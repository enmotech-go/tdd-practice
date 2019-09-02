package class04

import (
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got '%d' want '%d' given, '%v'", got, want, numbers)
	}
}
