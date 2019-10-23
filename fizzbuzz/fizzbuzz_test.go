package fizzbuzz

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T)  {
	t.Run("test_should_show_raw_number", func(t *testing.T) {
		fb:=FizzBuzz{1}
		assert.Equal(t,"1",fb.Number(1))
		fb=FizzBuzz{2}
		assert.Equal(t,"2",fb.Number(2))
	})
	t.Run("test_should_show_fizz", func(t *testing.T) {
		fb := FizzBuzz{3}
		assert.Equal(t,"Fizz",fb.Number(3))
	})
	t.Run("test_should_show_buzz", func(t *testing.T) {
		fb:=FizzBuzz{5}
		assert.Equal(t,"Buzz",fb.Number(5))
	})
	t.Run("test_should_show_fizz_buzz", func(t *testing.T) {
		fb:=FizzBuzz{15}
		assert.Equal(t,"FizzBuzz",fb.Number(15))
	})
	t.Run("test_should_show_fizz_or_buzz", func(t *testing.T) {
		fb:=FizzBuzz{13}
		assert.Equal(t,"Fizz",fb.Number(13))
		fb= FizzBuzz{55}
		assert.Equal(t,"Buzz",fb.Number(55))
		fb= FizzBuzz{51}
		assert.Equal(t,"FizzBuzz",fb.Number(51))
	})

}