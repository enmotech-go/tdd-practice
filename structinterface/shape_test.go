package structinterface

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{12.0, 5.0}
	got := rectangle.Perimeter()
	want := 34.0
	if got != want {
		t.Errorf("got %.2f but %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		Name  string
		Shape Shape
		Want  float64
	}{
		{"Rectangle", Rectangle{13.0, 4.0}, 52.0},
		{"Circle", Circle{21.0}, math.Pi * 21.0 * 21.0},
		{"Triangle", Triangle{12.0, 6.0}, 36.0},
	}

	for _, ss := range areaTests {
		t.Run(ss.Name, func(t *testing.T) {
			got := ss.Shape.Area()
			if math.Abs(got-ss.Want) > 0.0000001 {
				t.Errorf("%#v got %v, want %v", ss.Shape, got, ss.Want)
			}
		})
	}
}
