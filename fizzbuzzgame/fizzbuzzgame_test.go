package fizzbuzzgame

import "testing"
func TestFizzBuzzGame(t *testing.T)  {
	assert := func(t *testing.T, want, got string) {
		t.Helper()
		if want != got {
			t.Errorf("want %s , but Got %s", want, got)
		}
	}

	assert(t, FizzBuzzGame(1), "1")
	assert(t, FizzBuzzGame(2), "2")
	assert(t, FizzBuzzGame(3), "fizz")
	assert(t, FizzBuzzGame(5), "buzz")
	assert(t, FizzBuzzGame(15), "fizzbuzz")
	assert(t, FizzBuzzGame(13), "fizz")
	assert(t, FizzBuzzGame(51), "fizzbuzz")
	assert(t, FizzBuzzGame(35), "fizzbuzz")
}
