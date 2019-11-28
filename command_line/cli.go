package command_line
import (
	"io"
	"bufio"
	"strings"
	"os"
	"fmt"
)

type CLI struct {
	playerStore PlayerStore
	// in io.Reader

	in *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in: bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker()  {
	// cli.playerStore.RecordWin("Chris")

	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string  {
	return strings.Replace(userInput, " wins", "", 1)
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemStore, error)  {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", path, err)
	}
	store, err := NewFileSystemStore(db)
	if err != nil {
		return nil, fmt.Errorf("problem creating file system player store, %v", err)

	}

	return store, nil
}
