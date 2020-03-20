package string

import "testing"

type TestData struct {
	pattern string
	expect  int
}

func TestStrStrHorspool(t *testing.T) {
	target := "abcdefghijklmn"
	testData := []TestData{
		{pattern: "a", expect: 0},
		{pattern: "abc", expect: 0},
		{pattern: "cde", expect: 2},
		{pattern: "lmn", expect: 11},
		{pattern: "n", expect: 13},
		{pattern: "x", expect: -1},
		{pattern: "zx", expect: -1},
	}

	for _, pattern := range testData {
		ans := StrStrHorspool(target, pattern.pattern)
		if ans != pattern.expect {
			t.Logf("pattern: %s expect: %d ans: %d", pattern.pattern, pattern.expect, ans)
			t.FailNow()
		}
	}
}
