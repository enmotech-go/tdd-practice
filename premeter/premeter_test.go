package premeter

import "testing"

func TestPerimeter(t *testing.T) {
	got := Permeter(5.0, 10.0)
	want := 30.0

	if got != want {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 5.0)
	want := 60.0

	if got != want {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}
