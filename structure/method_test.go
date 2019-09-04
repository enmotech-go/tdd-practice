package structure

import (
	"math"
	"testing"
)

func TestPerimeter3(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	t.Run("test calculate perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 10.1, Height: 6.9}
		checkPerimeter(t, rectangle, 34.0)
	})

	t.Run("test calculate perimeter of circle", func(t *testing.T) {
		circle := Circle{Radius: 11.5}
		checkPerimeter(t, circle, math.Pi*11.5*2)
	})
}

func TestArea3(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	t.Run("test calculate area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 11.2, Height: 6.5}
		checkArea(t, rectangle, 72.8)
	})

	t.Run("test calculate area of circle", func(t *testing.T) {
		circle := Circle{Radius: 10}
		checkArea(t, circle, 314.1592653589793)
	})
}
