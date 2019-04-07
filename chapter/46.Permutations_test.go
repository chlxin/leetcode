package chapter

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}
