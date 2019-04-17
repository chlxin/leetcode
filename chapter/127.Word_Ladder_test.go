package chapter

import (
	"fmt"
	"testing"
)

func Test_ladderLength(t *testing.T) {
	res := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Println(res)
}

func Test_ladderLength2(t *testing.T) {
	res := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
	fmt.Println(res)
}

func Test_ladderLength3(t *testing.T) {
	res := ladderLength("hog", "cog", []string{"cog"})
	fmt.Println(res)
}
