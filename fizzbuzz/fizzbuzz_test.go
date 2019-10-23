package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("test should return raw number", func(t *testing.T) {
		result := FizzBuzz("1")
		expect := "1"
		if result != expect {
			t.Errorf("result(%s) != expect(%s)", result, expect)
		}
	})
}
