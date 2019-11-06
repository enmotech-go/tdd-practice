package ref

import (
	"reflect"
	"testing"
)

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
