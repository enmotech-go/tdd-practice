package structinterface

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(12.0, 5.0)
	want := 34.0
	if got != want {
		t.Errorf("got %.2f but %.2f", got, want)
	}
}
