package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("test_should_show_raw_number", func(t *testing.T) {
		fb := FizzBuzz{1}
		assert.Equal(t, "1", fb.Number())

		fb = FizzBuzz{2}
		assert.Equal(t, "2", fb.Number())
	})

	t.Run("test_should_show_fizz", func(t *testing.T) {
		fb := FizzBuzz{3}
		assert.Equal(t, "Fizz", fb.Number())
	})

	t.Run("test_should_show_buzz", func(t *testing.T) {
		fb := FizzBuzz{5}
		assert.Equal(t, "Buzz", fb.Number())
	})

	t.Run("test_should_show_fizzbuzz", func(t *testing.T) {
		fb := FizzBuzz{15}
		assert.Equal(t, "FizzBuzz", fb.Number())
	})
}
