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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 5.0}, 50.0},
		{Circle{10}, math.Pi * 10 * 10},
	}

	for _, ss := range areaTests {
		got := ss.shape.Area()
		if math.Abs(got-ss.want) > 0.00001 {
			t.Errorf("%#v, want %.2f, got %.2f", ss.shape, ss.want, got)
		}
	}

}
