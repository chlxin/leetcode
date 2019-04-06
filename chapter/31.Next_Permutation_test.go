package chapter

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
