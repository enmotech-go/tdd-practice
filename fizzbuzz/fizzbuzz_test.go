package fizzbuzz

import "testing"

func TestFizzBuzz_Convert(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		f    FizzBuzz
		args args
		want string
	}{
		{
			"test_3",
			FizzBuzz{},
			args{n: 3},
			WordFizz,
		},
		{
			"test_5",
			FizzBuzz{},
			args{n: 5},
			WordBuzz,
		},
		{
			"test_15",
			FizzBuzz{},
			args{n: 15},
			WordFizz + WordBuzz,
		},
		{
			"test_others",
			FizzBuzz{},
			args{n: 4},
			"4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Convert(tt.args.n); got != tt.want {
				t.Errorf("FizzBuzz.Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
