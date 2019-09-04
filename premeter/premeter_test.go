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

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 5.0}
		got := rectangle.Area()
		want := 60.0

		if got != want {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := math.Pi * circle.radius * circle.radius

		if got != want {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	})

}
