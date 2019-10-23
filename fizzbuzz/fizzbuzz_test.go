package fizzbuzz

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T)  {
	t.Run("test_should_show_raw_number", func(t *testing.T) {
		fb:=FizzBuzz{"1"}
		assert.Equal(t,"1",fb.Number())
		fb=FizzBuzz{"2"}
		assert.Equal(t,"2",fb.Number())
	})
}