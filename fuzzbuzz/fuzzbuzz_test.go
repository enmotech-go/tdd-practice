package fuzzbuzz

import "testing"

func TestFuzzBuzz(t *testing.T) {

	t.Run("saying fizz buzz", func(t *testing.T) {
		result := FizzBuzz("1")
		expect := "1"
		if result != expect {
			t.Errorf("result(%s) != expect(%s)",result,expect)
		}
	})

}
