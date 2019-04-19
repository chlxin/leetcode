package chapter

import (
	"fmt"
	"testing"
)

func Test_wordBreak2(t *testing.T) {
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	res := wordBreak2("catsanddog", wordDict)
	fmt.Println(res)
}

func Test_wordBreak21(t *testing.T) {
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	res := wordBreak2("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", wordDict)
	fmt.Println(res)
}

func Test_wordBreak3(t *testing.T) {
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	res := wordBreak3("catsanddog", wordDict)
	fmt.Println(res)
}

func Test_wordBreak31(t *testing.T) {
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	res := wordBreak3("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", wordDict)
	fmt.Println(res)
}
