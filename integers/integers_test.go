package integers

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {
	sum := Add(2, 3)
	want := 5

	if sum != want {
		t.Errorf("error want %d, but get %d", sum, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
