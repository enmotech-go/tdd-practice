package class11

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("test shloud show row number", func(t *testing.T) {
		got := FizzBuzz(1)
		assert.Equal(t, "1", got.Number())
	})
}
