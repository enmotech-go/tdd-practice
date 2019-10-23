package fizzbuzz

import "testing"

func Test_canBeDivisibleByThree(t *testing.T) {
	type args struct {
		set   int
		input int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "three can be divisible by three", args: args{set: 3, input: 3}, want: true},
		{name: "four can not be divisible by five", args: args{set: 5, input: 4}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeDivisibleBySet(tt.args.set, tt.args.input); got != tt.want {
				t.Errorf("canBeDivisibleBySet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSayText(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name       string
		args       args
		wantOutPut string
	}{
		// TODO: Add test cases.
		{name: "when set one out print one", args: args{input: 1}, wantOutPut: "1"},
		{name: "when set three out print Fizz", args: args{input: 3}, wantOutPut: Fizz},
		{name: "when set five out print Buzz", args: args{input: 5}, wantOutPut: Buzz},
		{name: "when set 15 out print FizzBuzz", args: args{input: 15}, wantOutPut: FizzBuzz},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutPut := GetSayText(tt.args.input); gotOutPut != tt.wantOutPut {
				t.Errorf("GetSayText() = %v, want %v", gotOutPut, tt.wantOutPut)
			}
		})
	}
}
