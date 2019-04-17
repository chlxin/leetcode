package chapter

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	res := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(res)
}

func Test_isPalindrome2(t *testing.T) {
	res := isPalindrome(".,")
	fmt.Println(res)
}

func Test_isPalindrome3(t *testing.T) {
	res := isPalindrome("0P")
	fmt.Println(res)
}
