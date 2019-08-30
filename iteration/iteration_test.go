package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("str")
	expected := "str__"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}
