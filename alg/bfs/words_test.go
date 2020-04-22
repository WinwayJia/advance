package bfs

import (
	"testing"
)

type testDataStruct struct {
	begin    string
	end      string
	list     []string
	expected int
}

func TestLadderLength(t *testing.T) {
	testData := []testDataStruct{
		{
			begin: "hit",
			end:   "cog",
			list:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
		},
		{
			begin: "a",
			end:   "c",
			list:  []string{"a", "b", "c"},
		},
	}

	for idx, item := range testData {
		ans := ladderLength(item.begin, item.end, item.list)
		if ans != 5 {
			t.Logf("index %d wrong answer %d", idx, ans)
			t.FailNow()
		}
	}
}
