package struct_method_interface

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("area of rectangle", func(t *testing.T) {
		got := Rectangle{10.0, 10.0}.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("area of circle", func(t *testing.T) {
		got := Circle{10.0}.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
