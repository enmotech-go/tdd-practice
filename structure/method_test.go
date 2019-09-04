package structure

import (
	"math"
	"testing"
)

func TestPerimeter3(t *testing.T) {
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

func TestArea3(t *testing.T) {
	t.Run("test calculate area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 11.2, Height: 6.5}
		got := rectangle.Area()
		want := 72.8

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("test calculate area of circle", func(t *testing.T) {
		circle := Circle{Radius: 10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %f, want %f", got, want)
		}
	})
}
