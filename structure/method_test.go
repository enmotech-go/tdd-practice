package structure

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("test calculate perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 10.1, Height: 6.9}
		got := rectangle.Perimeter()
		want := 34.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("test calculate perimeter of circle", func(t *testing.T) {
		circle := Circle{Radius: 11.5}
		got := circle.Perimeter()
		want := math.Pi * 11.5 * 2

		if got != want {
			t.Errorf("got %2.f, want %.2f", got, want)
		}
	})
}
