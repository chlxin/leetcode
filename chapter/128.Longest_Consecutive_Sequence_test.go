package chapter

import (
	"fmt"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(nums)
	fmt.Println(res)
}
