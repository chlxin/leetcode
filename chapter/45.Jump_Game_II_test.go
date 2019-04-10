package chapter

import (
	"fmt"
	"testing"
)

func Test_jump(t *testing.T) {
	// nums := []int{2, 3, 0, 1, 4}
	nums := []int{2, 0, 2, 4, 6, 0, 0, 3}
	// nums := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	res := jump(nums)
	fmt.Println(res)
}
