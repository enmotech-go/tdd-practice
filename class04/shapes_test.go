package class04

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %0.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 5.0)
	want := 60.0

	if got != want {
		t.Errorf("got %0.2f want %.2f", got, want)
	}
}
