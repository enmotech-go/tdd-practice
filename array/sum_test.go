package array

import (
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test sum 5 numbers",
			args{numbers: []int{1, 2, 3, 4, 5}},
			15,
		},
		{
			"test sum any size",
			args{numbers: []int{1, 2, 3}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers); got != tt.want {
				t.Errorf("SumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
