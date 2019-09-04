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
	cases := []struct {
		desc  string
		shape Shape
		want  float64
	}{
		{"test area fo rectangle", Rectangle{10.0, 10.0}, 100.0},
		{"test area of circle", Circle{10.0}, 314.1592653589793},
		{"test area of triangle", Triangle{12, 6}, 36.0},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := c.shape.Area()
			if got != c.want {
				t.Errorf("%s got %.2f want %.2f", c.desc, got, c.want)
			}
		})
	}
}
