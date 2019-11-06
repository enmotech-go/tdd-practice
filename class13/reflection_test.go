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
			Person{
				Name: "Chris",
				Profile: Profile{
					33,
					"London",
				},
			},
			[]string{"Chris", "London"},
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

}
