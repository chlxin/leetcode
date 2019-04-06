package chapter

import (
	"fmt"
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	res := longestValidParentheses(")()())")
	fmt.Println(res)
}
