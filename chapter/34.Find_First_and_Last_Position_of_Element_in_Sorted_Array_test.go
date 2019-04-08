package chapter

import (
	"fmt"
	"testing"
)

func Test_searchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	res := searchRange(nums, 6)
	fmt.Println(res)
}

func Test_searchRange_Empty(t *testing.T) {
	nums := []int{}
	res := searchRange(nums, 0)
	fmt.Println(res)
}
