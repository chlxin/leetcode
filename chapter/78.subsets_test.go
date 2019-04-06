package chapter

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}

func Test_subsetsDP(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsetsDP(nums)
	fmt.Println(res)
}
