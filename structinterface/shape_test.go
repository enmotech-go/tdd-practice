package structinterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{12.0, 5.0}
	got := Perimeter(rectangle)
	want := 34.0
	if got != want {
		t.Errorf("got %.2f but %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 5.0}
	got := Area(rectangle)
	want := 60.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
