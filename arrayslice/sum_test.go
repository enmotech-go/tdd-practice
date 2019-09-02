package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test array", func(t *testing.T) {
		number := []int{1, 2, 3, 4, 5}
		got := Sum(number)
		want := 15
		if want != got {
			t.Errorf("got '%d', but want '%d' given %v", got, want, number)
		}
	})

	t.Run("test slice", func(t *testing.T) {
		number := []int{2, 3, 4, 5, 6}
		got := Sum(number)
		want := 20
		if want != got {
			t.Errorf("got %d want %d give %v", got, want, number)
		}
	})
}

func TestSumall(t *testing.T) {
	numberList := [][]int{{1, 2, 3}, {4, 5, 6}}
	got := sumAll(numberList)
	want := []int{6, 15}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v, numberList %v", got, want, numberList)
	}
}

func TestSumAllTails(t *testing.T) {
	numberList := [][]int{{1, 2, 3}, {4, 5, 6}}
	got := sumAllTails(numberList)
	want := []int{5, 11}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v, numberList %v", got, want, numberList)
	}
}
