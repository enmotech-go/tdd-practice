package class05

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Perimeter() = %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 8.0)
	want := 96.0

	if got != want {
		t.Errorf("Area() = %.2f, want %.2f", got, want)
	}
}
