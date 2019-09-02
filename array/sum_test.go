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
		numbers [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test sum two slices",
			args{
				numbers: [][]int{{1, 2}, {0, 9}},
			},
			[]int{3, 9},
		},
		{
			"test sum three slices",
			args{
				numbers: [][]int{{1, 2}, {0, 9}, {4, 4}},
			},
			[]int{3, 9, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAll(tt.args.numbers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumAllTails(t *testing.T) {
	type args struct {
		numbers [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test sum two slices tails",
			args{
				numbers: [][]int{{1, 2, 3}, {0, 9}},
			},
			[]int{5, 9},
		},
		{
			"test sum empty slice tails",
			args{
				numbers: [][]int{{1, 2, 3}, {}},
			},
			[]int{5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAllTails(tt.args.numbers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAllTails() = %v, want %v", got, tt.want)
			}
		})
	}
}
