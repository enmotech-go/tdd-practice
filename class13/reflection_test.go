package class13

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})
	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d wnat %d ", len(got), 1)
	}
}
