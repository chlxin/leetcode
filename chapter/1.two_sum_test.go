package chapter

import "testing"

func TestTwoSum(t *testing.T)  {
	ints := twoSum([]int{2, 7, 11, 15}, 9)
	if ints[0] != 0 || ints[1] != 1 {
		t.Fail()
	}
}
