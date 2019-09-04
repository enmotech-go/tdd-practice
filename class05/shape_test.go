package class05

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Perimeter() = %.2f, want %.2f", got, want)
	}
}
