package iteration

import "testing"

func TestRepeat(t*testing.T){
	repeated:=Repeat("a")
	expected:="aaaaa"

	if repeated != expected{
		t.Errorf("expected %s bug got %s", expected, repeated)
	}
}