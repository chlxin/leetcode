package chapter

import (
	"fmt"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	res := wordBreak("leetcode", []string{"leet", "code"})
	fmt.Println(res)
}
