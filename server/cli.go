package poker

type CLI struct {
	playerStore PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Cleo")
}
