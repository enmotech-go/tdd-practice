package fizzbuzz

import "testing"

func assertEqual(t *testing.T, result, expect string) {
	t.Helper()
	if result != expect {
		t.Errorf("result(%s) != expect(%s)", result, expect)
	}
}

func TestFizzBuzz(t *testing.T) {
	t.Run("test should return raw number", func(t *testing.T) {
		assertEqual(t, FizzBuzz(1), "1")
	})

	t.Run("test should return fizz", func(t *testing.T) {
		assertEqual(t, FizzBuzz(3), "fizz")
	})

	t.Run("test should return buzz", func(t *testing.T) {
		assertEqual(t, FizzBuzz(5), "buzz")
	})

	t.Run("test should return fizzbuzz", func(t *testing.T) {
		assertEqual(t, FizzBuzz(15), "fizzbuzz")
	})

	t.Run("test should return fizzbuzz with 51", func(t *testing.T) {
		assertEqual(t, FizzBuzz(51), "fizzbuzz")
	})
}
