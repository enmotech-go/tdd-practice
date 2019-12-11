package poker_test

import (
	poker "baikai/tdd-practice/server"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	//in := strings.NewReader("Chris wins\n")
	//playerStore := &StubPlayerStore{}
	//cli := &CLI{playerStore, in}
	//cli.PlayPoker()
	//assertPlayerWin(t,playerStore,"Chris")
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
