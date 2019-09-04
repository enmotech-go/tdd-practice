package structure

import "testing"

func TestPremeter(t *testing.T) {
	type args struct {
		width  float64
		height float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test calculate premeter",
			args{
				width:  10.1,
				height: 6.9,
			},
			34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Premeter(tt.args.width, tt.args.height); got != tt.want {
				t.Errorf("Premeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
