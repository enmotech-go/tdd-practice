package poker

import "io"

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Chris")
}
