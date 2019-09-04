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

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{12.0, 5.0}, want: 60.00},
		{name: "circles", shape: Circle{R: 10.0}, want: math.Pi * 10.0 * 10.0},
		{name: "circles", shape: Triangle{12.0, 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if tt.want != got {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})

	}
}
