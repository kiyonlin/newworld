package poker

import "testing"

// StubPlayerStore implements PlayerStore for testing purposes
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

// AssertPlayerWin allows you to spy on the store's calls to RecordWin
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("实际调用了 RecordWin %d 次 期望 %d 次", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("没有保存正确的胜利者 结果 '%s' 期望 '%s'", store.WinCalls[0], winner)
	}
}
