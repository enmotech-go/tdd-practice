package premeter

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.0, 10.0}
	got := rectangle.Permeter()
	want := 30.0

	if got != want {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 5.0}
		checkArea(t, rectangle, 60.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, math.Pi*circle.radius*circle.radius)
	})
}
