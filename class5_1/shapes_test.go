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
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			W: 12.0,
			H: 5.0,
		}
		want := 60.0
		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := math.Pi * 10.0 * 10.0
		checkArea(t, circle, want)
	})

}
