package fuzzbuzz

import "testing"

func TestFuzzBuzz(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("saying number", func(t *testing.T) {
		result := FizzBuzz(1)
		got := "1"
		assertCorrectMessage(t,result,got)
	})



	t.Run("get fizz",func(t *testing.T){
		result := FizzBuzz(3)
		got := Fizz
		assertCorrectMessage(t,result,got)
	})

	t.Run("saying fizz", func(t *testing.T) {
		result := FizzBuzz(13)
		got := Fizz
		assertCorrectMessage(t,result,got)
	})

	t.Run("get buzz",func(t *testing.T){
		result := FizzBuzz(5)
		got := Buzz
		assertCorrectMessage(t,result,got)
	})
	t.Run("get fizzbuzz",func(t *testing.T){
		result := FizzBuzz(15)
		got := FizzBuzzStr
		assertCorrectMessage(t,result,got)
	})
}
