package class11

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("test should show row number", func(t *testing.T) {
		got := FizzBuzz(1)
		assert.Equal(t, "1", got.Number())
		got = FizzBuzz(2)
		assert.Equal(t, "2", got.Number())
	})
	t.Run("test should show row 3 number show Fizz", func(t *testing.T) {
		got := FizzBuzz(3)
		assert.Equal(t, "Fizz", got.Number())

	})

	t.Run("test should show row 5 number show Buzz", func(t *testing.T) {
		got := FizzBuzz(5)
		assert.Equal(t, "Buzz", got.Number())

	})

	t.Run("test should show row 3 and 5 number show Buzz", func(t *testing.T) {
		got := FizzBuzz(15)
		assert.Equal(t, "FizzBuzz", got.Number())

	})

	t.Run("test should show row 3  number show Fizz", func(t *testing.T) {
		got := FizzBuzz(13)
		assert.Equal(t, "Fizz", got.Number())

	})
	t.Run("test should show row 3  number show FizzBuzz", func(t *testing.T) {
		got := FizzBuzz(51)
		assert.Equal(t, "FizzBuzz", got.Number())

	})
}
