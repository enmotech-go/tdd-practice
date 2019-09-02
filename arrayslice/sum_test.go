package sum

import "testing"

func TestSum(t *testing.T) {
	number := [5]int{1, 2, 3, 4, 5}
	got := Sum(number)
	want := 15
	if want != got {
		t.Errorf("got '%d', but want '%d' given %v", got, want, number)
	}
}
