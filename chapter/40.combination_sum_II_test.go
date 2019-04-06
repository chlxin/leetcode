package chapter

import (
	"fmt"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	// nums := []int{10, 1, 2, 7, 6, 1, 5}
	nums := []int{2, 5, 2, 1, 2}
	res := combinationSum2(nums, 5)
	fmt.Println(res)
}
