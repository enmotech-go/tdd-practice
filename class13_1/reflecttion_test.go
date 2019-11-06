package class13_1

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"Struct with one string field",
			&Person{"Chris", Profile{
				Age:  33,
				City: "London",
			}},
			[]string{"Chris", "London"},
		},
		{"Struct with one string field",
			[]Profile{
				{
					Age:  33,
					City: "London",
				}, {Age: 34, City: "Reykjavík"}},
			[]string{"London", "Reykjavík"},
		},

		{"Arrays",
			[2]Profile{
				{
					Age:  33,
					City: "London",
				}, {Age: 34, City: "Reykjavík"}},
			[]string{"London", "Reykjavík"},
		},
		{"Maps",
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
			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tt.ExpectedCalls)
			}
		})
	}

	t.Run("with Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

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
