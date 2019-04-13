package chapter

import (
	"fmt"
	"testing"
)

func Test_canJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	res := canJump(nums)
	fmt.Println(res)
}

func Test_canJump2(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	res := canJump(nums)
	fmt.Println(res)
}

func Test_canJump3(t *testing.T) {
	nums := []int{0}
	res := canJump(nums)
	fmt.Println(res)
}
