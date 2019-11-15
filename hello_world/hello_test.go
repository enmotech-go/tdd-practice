package main

import (
	"testing"
)

// func TestHello(t *testing.T) {
// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Chris", "Chris")
// 		want := "Hello, Chris"
// 		assertEqual(t, got, want)
// 	})
// 	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
// 		got := Hello("", "")
// 		want := "Hello, world"
// 		assertEqual(t, got, want)
// 	})
// 	t.Run("in Spanish", func(t *testing.T) {
// 		got := Hello("Elodie", "Spanish")
// 		want := "Hola, Elodie"
// 		assertEqual(t, got, want)
// 	})
// 	t.Run("in French", func(t *testing.T) {
// 		got := Hello("fman", "French")
// 		want := "FHello, fman"
// 		assertEqual(t, got, want)
// 	})

// }

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}
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
			args{"", "Spanish"},
			"Hola, world",
		},
		{
			"say hello to Mary in $%?@!",
			args{"Mary", "$%?@!"},
			"Hello, Mary",
		},
		{
			"say hello to ff in French",
			args{"ff", "French"},
			"FHello, ff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertEqual(t, Hello(tt.args.name, tt.args.lang), tt.want)
			// if got := Hello(tt.args.name, tt.args.lang); got != tt.want {
			// 	t.Errorf("Hello() = %v, want %v", got, tt.want)
			// }
		})
	}
}
