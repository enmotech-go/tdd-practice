package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("test_should_show_raw_number", func(t *testing.T) {
		assert.Equal(t, "1", FizzBuzz{1}.Number())
		assert.Equal(t, "2", FizzBuzz{2}.Number())
	})

	t.Run("test_should_show_fizz", func(t *testing.T) {
		assert.Equal(t, "Fizz", FizzBuzz{3}.Number())
	})

	t.Run("test_should_show_buzz", func(t *testing.T) {
		assert.Equal(t, "Buzz", FizzBuzz{5}.Number())
	})

	t.Run("test_should_show_fizzbuzz", func(t *testing.T) {
		assert.Equal(t, "FizzBuzz", FizzBuzz{15}.Number())
	})

	t.Run("test_should_show_fizz_or_buzz", func(t *testing.T) {
		assert.Equal(t, "Fizz", FizzBuzz{13}.Number())
		assert.Equal(t, "FizzBuzz", FizzBuzz{51}.Number())
		assert.Equal(t, "Buzz", FizzBuzz{52}.Number())
		assert.Equal(t, "FizzBuzz", FizzBuzz{35}.Number())
	})
}
