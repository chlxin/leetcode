package chapter

import (
	"fmt"
	"testing"
)

func Test_getPermutation(t *testing.T) {
	res := getPermutation(4, 9)
	fmt.Println(res)
}
