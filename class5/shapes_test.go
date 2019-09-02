package class5

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	if got != want {
		t.Errorf("got '%f' want '%f'", got, want)
	}
}
