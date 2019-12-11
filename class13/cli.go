package poker

import "io"

type CLI struct {
	PlayerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.PlayerStore.RecordWin("Chris")
}
