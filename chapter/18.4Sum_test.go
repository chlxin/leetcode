package chapter

import (
	"fmt"
	"testing"
)

func Test_fourSum(t *testing.T) {
	// [-1,-5,-5,-3,2,5,0,4]
	// -7
	// nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{-1, -5, -5, -3, 2, 5, 0, 4}
	res := fourSum(nums, -7)
	fmt.Println(res)
}
