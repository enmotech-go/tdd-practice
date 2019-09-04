package class5_1

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(12.0, 5.0)
	want := 34.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{
		W: 12.0,
		H: 5.0,
	}
	got := Area(rectangle)
	want := 60.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
