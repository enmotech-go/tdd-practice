package class5

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0
	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		dim := math.Dim(got, want)
		if dim < 0.0000001 {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.00
		got := rectangle.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	// flaot 因为底层存放的问题，并不是一个准确的值，所以在比较的时候不能直接进行相等比较，而在使用精度比较的时候，设置精度和比较位数一样，如果使用第二种比比较为更精确一位则两个数就不相等了。
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.159265
		checkArea(t, &circle, want)

	})
}
