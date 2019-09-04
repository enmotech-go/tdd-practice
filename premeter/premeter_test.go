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
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 5.0}, want: 50.0},
		{name: "Circle", shape: Circle{10}, want: math.Pi * 10 * 10},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, want: 36.0},
	}

	for _, ss := range areaTests {

		t.Run(ss.name, func(t *testing.T) {
			got := ss.shape.Area()
			if math.Abs(got-ss.want) > 0.00001 {
				t.Errorf("%#v, want %.2f, got %.2f", ss.shape, ss.want, got)
			}
		})
	}
}
