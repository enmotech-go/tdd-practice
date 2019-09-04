package class04

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %0.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{Width: 12.0, Height: 5.0}
	got := rectangle.Area()
	want := 60.0

	if got != want {
		t.Errorf("got %0.2f want %.2f", got, want)
	}
}
