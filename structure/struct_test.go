package structure

import "testing"

func TestPerimeter2(t *testing.T) {
	type args struct {
		rectangle Rectangle
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test calculate perimeter of rectangle",
			args{
				rectangle: Rectangle{Width: 10.1, Height: 6.9},
			},
			34.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perimeter2(tt.args.rectangle); got != tt.want {
				t.Errorf("Perimeter2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArea2(t *testing.T) {
	type args struct {
		rectangle Rectangle
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test calculate area of rectangle",
			args{
				rectangle: Rectangle{Width: 11.2, Height: 6.5},
			},
			72.8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Area2(tt.args.rectangle); got != tt.want {
				t.Errorf("Area2() = %v, want %v", got, tt.want)
			}
		})
	}
}
