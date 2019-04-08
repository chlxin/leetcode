package chapter

import (
	"fmt"
	"testing"
)

func Test_twoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	res := twoSum2(nums, 9)
	fmt.Println(res)
}
