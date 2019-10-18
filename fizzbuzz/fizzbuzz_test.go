package fizzbuzz

import (
	"bytes"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			"test_3",
			args{n: 3},
			`1
2
Fizz
`,
		},
		{
			"test_5",
			args{n: 5},
			`1
2
Fizz
4
Buzz
`,
		},
		{
			"test_15",
			args{n: 15},
			`1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			FizzBuzz(w, tt.args.n)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("FizzBuzz() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
