package class04

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %0.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("Rectangle area", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 5.0}
		got := rectangle.Area()
		want := 60.0

		if got != want {
			t.Errorf("got %0.2f want %.2f", got, want)
		}
	})

	t.Run("Circle area", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := math.Pi * 10 * 10

		if got != want {
			t.Errorf("got %0.2f want %.2f", got, want)
		}
	})
}
