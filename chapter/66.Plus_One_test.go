package chapter

import (
	"fmt"
	"testing"
)

func Test_plusOne(t *testing.T) {
	digits := []int{9}
	res := plusOne(digits)
	fmt.Println(res)
}
