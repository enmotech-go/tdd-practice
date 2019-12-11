package poker

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {
	in := strings.NewReader("Chris wins\n")
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore, in}
	cli.PlayPoker()

	assert.GreaterOrEqual(t, len(playerStore.winCalls), 1)

	got := playerStore.winCalls[0]
	want := "Chris"
	assert.Equal(t, want, got)
}
