package class5_1

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		W: 12.0,
		H: 5.0,
	}
	got := Perimeter(rectangle)
	want := 34.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			W: 12.0,
			H: 5.0,
		}
		got := rectangle.Area()
		want := 60.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := Circle.Area(circle)
		want := math.Pi * 10.0 * 10.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
