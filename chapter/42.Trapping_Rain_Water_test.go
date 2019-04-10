package chapter

import (
	"fmt"
	"testing"
)

func Test_trap(t *testing.T) {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(nums)
	fmt.Println(res)
}
