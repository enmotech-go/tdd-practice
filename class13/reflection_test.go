package class13

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(input string) {
				got = append(got, input)
			})
			assert.Equal(t, got, tt.ExpectedCalls)
		})

	}

	t.Run("with map", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz"}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")

	})

}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}

	}
	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didnt", haystack, needle)
	}

}
