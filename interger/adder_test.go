package interger

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T)  {
	got := Add(2, 2)
	want := 4
	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}

func ExampleAdd()  {
	sum := Add(1 ,2)
	fmt.Println("sum: ", sum)
}