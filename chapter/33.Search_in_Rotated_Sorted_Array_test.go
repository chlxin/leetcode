package chapter

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	// nums := []int{5, 1, 3}
	res := search(nums, 5)
	fmt.Println(res)
}
