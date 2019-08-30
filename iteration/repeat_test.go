package iteration

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test repeat empty char",
			args{char: ""},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.char); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
