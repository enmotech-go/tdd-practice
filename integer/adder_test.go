package integers

import "testing"
import "fmt"

func TestAddr(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d ,got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	//Output:3
}
