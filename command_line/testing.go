package command_line
import (
	"testing"
)

// 存根
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string // 监视收集
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// AssertPlayerWin allows you to spy on the store's calls to RecordWin
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string)  {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got is %d call to RecordWin want is %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store the correct winner got is %q, want is %q", store.winCalls[0], winner)
	}
}

