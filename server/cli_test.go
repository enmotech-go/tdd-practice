package poker

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()
	assert.Len(t, store.winCalls, 1)
	assert.Equal(t, winner, store.winCalls[0])
}

func TestCli(t *testing.T) {
	t.Run("record_chris_win_from_user_input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}
		cli := NewCLI(playerStore, in)
		cli.PlayPoker()
		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record_cleo_win_from_user_input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}
		cli := NewCLI(playerStore, in)
		cli.PlayPoker()
		assertPlayerWin(t, playerStore, "Cleo")
	})
}
