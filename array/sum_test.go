package array

import (
	"testing"
)

func TestSumArray(t *testing.T) {
	type args struct {
		numbers [5]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test sum array",
			args{numbers: [5]int{1, 2, 3, 4, 5}},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumArray(tt.args.numbers); got != tt.want {
				t.Errorf("SumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSlice(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test sum empty slice",
			args{numbers: []int{}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSlice(tt.args.numbers); got != tt.want {
				t.Errorf("SumSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
