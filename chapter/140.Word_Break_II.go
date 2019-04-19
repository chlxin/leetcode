package chapter

import "strings"

// 太慢了，TLE
func wordBreak2(s string, wordDict []string) []string {
	str := []byte(s)
	words := make(map[string]bool)
	for _, word := range wordDict {
		words[word] = true
	}
	res := make([][]string, 0)
	backtrack140(&res, str, make([]string, 0), words, 0)
	result := make([]string, 0)
	for _, solve := range res {
		t := strings.Join(solve, " ")
		result = append(result, t)
	}
	return result
}

func backtrack140(res *[][]string, str []byte, curList []string, words map[string]bool, start int) {
	if len(curList) > 0 && start >= len(str) {
		tmpList := make([]string, len(curList))
		copy(tmpList, curList)
		*res = append(*res, tmpList)
		return
	}

	for i := start; i <= len(str); i++ {
		if words[string(str[start:i])] {
			curList = append(curList, string(str[start:i]))
			backtrack140(res, str, curList, words, i)
			curList = curList[:len(curList)-1]
		}
	}
}

func wordBreak3(s string, wordDict []string) []string {
	words := make(map[string]bool)
	for _, word := range wordDict {
		words[word] = true
	}
	memory := make(map[string][]string)
	return DFS(s, words, memory)
}

// DFS 搜索
func DFS(s string, words map[string]bool, memory map[string][]string) []string {
	if v, ok := memory[s]; ok {
		return v
	}

	res := make([]string, 0)
	if len(s) == 0 {
		res = append(res, "") // 需要增加一个空的字符串，这样后面的subList for range时候才会有元素遍历，才会往res里写东西
		return res
	}

	for word := range words {
		if strings.HasPrefix(s, word) {
			subList := DFS(s[len(word):], words, memory)
			for _, sub := range subList {
				if sub == "" {
					res = append(res, word)
				} else {
					res = append(res, word+" "+sub)
				}
			}
		}
	}
	memory[s] = res
	return res
}
