package chapter

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool)
	visited := make(map[string]bool)
	for _, word := range wordList {
		wordMap[word] = true
	}
	if !wordMap[endWord] {
		return 0
	}

	var (
		word     string
		distance int
	)

	q := make([]string, 0, len(wordList))
	q = append(q, beginWord)
	visited[beginWord] = true
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			word, q = q[0], q[1:]
			for i := 0; i < len(word); i++ {
				bs := []byte(word)
				for ch := 'a'; ch <= 'z'; ch++ {
					bs[i] = byte(ch)
					newWord := string(bs)
					if newWord == endWord {
						return distance + 2
					}
					if wordMap[newWord] && !visited[newWord] {
						q = append(q, newWord)
						visited[newWord] = true
					}
				}
			}
		}
		distance++
	}

	return 0
}
