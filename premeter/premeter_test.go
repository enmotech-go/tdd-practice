package premeter

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.0, 10.0}
	got := Permeter(rectangle)
	want := 30.0

	if got != want {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 5.0}
	got := Area(rectangle)
	want := 60.0

	if got != want {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}
