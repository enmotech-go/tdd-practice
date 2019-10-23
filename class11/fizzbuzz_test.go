package class11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzz(t *testing.T) {
	cases := []struct {
		In   int
		Wnat string
	}{
		{In: 1, Wnat: "1"},
		{In: 0, Wnat: "0"},
		{In: 3, Wnat: "fizz"},
		{In: 4, Wnat: "4"},
		{In: 5, Wnat: "buzz"},
		{In: 15, Wnat: "fizzbuzz"},
	}
	for _, c := range cases {
		got := Fizzbuzz(c.In)
		assert.Equal(t, c.Wnat, got)
	}
}
