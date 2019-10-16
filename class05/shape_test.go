package class05

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("Perimeter() = %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{12, 8}, 96.0},
		{"circle", Circle{10.0}, math.Pi * 10.0 * 10.0},
		{"triangle", Triangle{12.0, 6.0}, 36.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("Area() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}
