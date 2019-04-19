package chapter

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	res := singleNumber(nums)
	fmt.Println(res)
}
