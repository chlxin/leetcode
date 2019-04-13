package chapter

import (
	"fmt"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	res := lengthOfLastWord("Hello World")
	fmt.Println(res)
}

func Test_lengthOfLastWord2(t *testing.T) {
	res := lengthOfLastWord("Hello World ")
	fmt.Println(res)
}

func Test_lengthOfLastWord3(t *testing.T) {
	res := lengthOfLastWord("b   a    ")
	fmt.Println(res)
}
