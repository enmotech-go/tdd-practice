package reflection

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
		{
			"Struct with one string field",
			struct{ Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Xiaoming"},
			},
			[]string{"London", "Xiaoming"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Xiaoming"},
			},
			[]string{"London", "Xiaoming"},
		},
	}

	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			var got []string
			walk(v.Input, func(input string) {
				got = append(got, input) // 元素追加
			})

			if !reflect.DeepEqual(got, v.ExpectedCalls) {
				t.Errorf("got %v, want %v.", got, v.ExpectedCalls)
			}
		})
	}
	t.Log("reflect finished")
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %s, but it didn't.", haystack, needle)
	}

}
