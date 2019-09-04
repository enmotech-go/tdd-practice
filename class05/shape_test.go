package class05

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("Perimeter() = %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		got := Rectangle{12.0, 8.0}.Area()
		want := 96.0

		if got != want {
			t.Errorf("Area() = %.2f, want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		got := Circle{10.0}.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("Area() = %.2f, want %.2f", got, want)
		}
	})
}
