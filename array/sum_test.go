package array

import (
	"reflect"
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

func TestSumAll(t *testing.T) {
	type args struct {
		numbers1 []int
		numbers2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test sum two slices",
			args{
				numbers1: []int{1, 2},
				numbers2: []int{0, 9},
			},
			[]int{3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAll(tt.args.numbers1, tt.args.numbers2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
