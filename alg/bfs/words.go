package bfs

// 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
//    每次转换只能改变一个字母。
//    转换过程中的中间单词必须是字典中的单词。
/* 说明:

   如果不存在这样的转换序列，返回 0。
   所有单词具有相同的长度。
   所有单词只由小写字母组成。
   字典中不存在重复的单词。
   你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
*/

func isNear(x, y string) bool {
	diff := 0
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	maps := make(map[string][]string, len(wordList))
	for i := 0; i < len(wordList)-1; i++ {
		for j := i + 1; j < len(wordList); j++ {
			if isNear(wordList[i], wordList[j]) {
				maps[wordList[i]] = append(maps[wordList[i]], wordList[j])
				maps[wordList[j]] = append(maps[wordList[j]], wordList[i])
			}
		}
	}

	type node struct {
		word   string
		length int
	}

	ans := 0
	for i := 0; i < len(wordList); i++ {
		if !isNear(beginWord, wordList[i]) {
			continue
		}

		queue := make([]node, 0, len(wordList))
		visited := make(map[string]bool, len(wordList))

		queue = append(queue, node{word: wordList[i], length: 1})
		visited[wordList[i]] = true

		for len(queue) != 0 {
			word := queue[0]
			queue = queue[1:]
			items, ok := maps[word.word]
			if !ok {
				continue
			}
			for _, item := range items {
				if item == endWord {
					if 0 == ans || ans > word.length+1+1 {
						ans = word.length + 1 + 1
					}
					break
				}
				_, ok = visited[item]
				if ok {
					continue
				}
				queue = append(queue, node{
					word:   item,
					length: word.length + 1,
				})
				visited[item] = true
			}
		}
	}

	return ans
}
