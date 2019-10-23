package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("test should show row number", func(t *testing.T) {
		fb := FizzBuzz{1}
		assert.Equal(t, "1", fb.Number())
	})

	t.Run("test num % 3 == 0", func(t *testing.T) {
		fb := FizzBuzz{3}
		assert.Equal(t, "fizz", fb.Number())
	})

	t.Run("test num % 5 == 0", func(t *testing.T) {
		fb := FizzBuzz{5}
		assert.Equal(t, "buzz", fb.Number())
	})

	t.Run("test num % 3||5 == 0", func(t *testing.T) {
		fb := FizzBuzz{15}
		assert.Equal(t, "fizzbuzz", fb.Number())
	})

	t.Run("test should show fizz or buzz", func(t *testing.T) {
		assert.Equal(t, "fizz", FizzBuzz{13}.Number())
		assert.Equal(t, "buzz", FizzBuzz{52}.Number())
		assert.Equal(t, "fizzbuzz", FizzBuzz{53}.Number())
		assert.Equal(t, "fizzbuzz", FizzBuzz{51}.Number())
	})
}
