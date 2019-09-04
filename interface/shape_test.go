package shape

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := math.Pi * 10 * 10

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
