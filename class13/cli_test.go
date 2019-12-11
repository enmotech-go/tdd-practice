package poker_test

import (
	"strings"
	"testing"

	poker "enmotech-go/tdd-practice/class13"
)

func TestCli(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cleo")
	})
}

func assertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any one")
	}
	if store.winCalls[0] != winner {
		t.Errorf("did not record correct winner, got '%s' ,want '%s'", store.winCalls[0], winner)
	}
}
