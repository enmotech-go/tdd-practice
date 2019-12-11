package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore}
	cli.PlayPoker()

	assert.Len(t, playerStore.winCalls, 1)
}
