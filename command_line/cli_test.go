package poker_test

import (
	"testing"
	"strings"
	"io"
	poker "enmotech-go/tdd-practice/command_line"

)

func TestCLI(t *testing.T)  {
	t.Run("record chris win from user input", func (t *testing.T)  {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		// playerStore := &poker.StubPlayerStore{}
		// cli := &poker.CLI{playerStore, in}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		// if len(playerStore.winCalls) < 1 {
		// 	t.Fatalf("expected a win call but didn't get any, got is %v", playerStore.winCalls)
		// }

		// got := playerStore.winCalls[0]
		want := "Chris"
		// assertResponseBody(t, got, want)

		// assertPlayerWin(t, playerStore, want)
		poker.AssertPlayerWin(t, playerStore, want)
	})

	t.Run("record cleo win from user input", func (t *testing.T)  {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		// cli := &poker.CLI{playerStore, in}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		want := "Cleo"
		// assertPlayerWin(t, playerStore, want)
		poker.AssertPlayerWin(t, playerStore, want)

	})

	t.Run("do not read beyond the first newline", func (t *testing.T)  {
		in := failOnEndReader{
			t, strings.NewReader("Chris wins\n hello there"),
		}
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
	})
}

type failOnEndReader struct{
	t *testing.T
	rdr io.Reader
}

func (m failOnEndReader) Read(p []byte) (n int, err error)  {
	n,err = m.rdr.Read(p)
	if n == 0 || err == io.EOF {
		m.t.Fatal("Read to the end when you shoudn't hava")
	}

	return n,err
}

// func assertPlayerWin(t *testing.T, store *poker.StubPlayerStore, winner string)  {
// 	t.Helper()
// 	if len(store.winCalls) != 1 {
// 		t.Fatalf("got is %d, calls to RecordWin want is %d", len(store.winCalls), 1)
// 	}
// 	if store.winCalls[0] != winner {
// 		t.Errorf("did not store correct winner got is '%s', want is '%s'", store.winCalls[0], winner)
		
// 	}

// }
