package poker_test

import (
	"strings"
	"testing"

	poker "tdd-practice/server"
)

func TestCli(t *testing.T) {
	t.Run("record_chris_win_from_user_input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record_cleo_win_from_user_input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
