package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "Chris")
		want := "Hello, Chris"
		assertEqual(t, got, want)
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertEqual(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertEqual(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("fman", "French")
		want := "FHello, fman"
		assertEqual(t, got, want)
	})

}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}
