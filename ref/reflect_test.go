package ref

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
			"test_struct_with_one_string_field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"test_struct_with_two_string_fields",
			struct {
				Name string
				City string
			}{"Chris", "New York"},
			[]string{"Chris", "New York"},
		},
		{
			"test_with_non_string_field",
			struct {
				Name string
				Age  int
			}{"Chris", 30},
			[]string{"Chris"},
		},
		{
			"test_with_nested_fields",
			Person{
				"Chris",
				Profile{30, "New York"},
			},
			[]string{"Chris", "New York"},
		},
		{
			"test_pointer",
			&Person{
				"Chris",
				Profile{30, "New York"},
			},
			[]string{"Chris", "New York"},
		},
		{
			"test_with_slices",
			[]Profile{
				{30, "New York"},
				{20, "London"},
			},
			[]string{"New York", "London"},
		},
		{
			"test_with_arrays",
			[2]Profile{
				{30, "New York"},
				{20, "London"},
			},
			[]string{"New York", "London"},
		},
		{
			"test_with_map",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("want %v, got %v", test.ExpectedCalls, got)
			}
		})
	}
}
