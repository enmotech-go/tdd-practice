package iteration

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		char  string
		times int
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
		{
			"test repeat `hello`",
			args{char: "hello", times: 2},
			"hellohello",
		},
		{
			"test repeat `world` 4 times",
			args{char: "world", times: 4},
			"worldworldworldworld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.char, tt.args.times); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
