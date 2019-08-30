package class2

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("got '%d' want '%d' ", got, want)
	}

}

func ExampleAdd() {
	add := Add(1, 5)
	fmt.Println(add)
	// Output: 6
}
