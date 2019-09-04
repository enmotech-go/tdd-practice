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
		shape Shape
		want  float64
	}{
		{Rectangle{12, 8}, 96.0},
		{Circle{10.0}, math.Pi * 10.0 * 10.0},
	}

	for _, tt := range tests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Area() = %v, want %v", got, tt.want)
		}
	}
}
