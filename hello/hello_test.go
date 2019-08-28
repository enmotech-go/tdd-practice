package hello

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name string
		lang string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"say hello to John in English",
			args{"John", "English"},
			"Hello, John",
		},
		{
			"say hello to anonymous in Chinese",
			args{"", "Chinese"},
			"你好，World",
		},
		{
			"say hello to Mary in $%?@!",
			args{"Mary", "$%?@!"},
			"Hello, Mary",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name, tt.args.lang); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
