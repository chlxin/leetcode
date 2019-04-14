package chapter

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	roundRes := make([]int, 0)
	sort.Ints(nums)
	backtrack90(&res, nums, roundRes, 0)
	return res
}

func backtrack90(res *[][]int, nums []int, roundRes []int, index int) {
	tmpRes := make([]int, len(roundRes))
	copy(tmpRes, roundRes)
	*res = append(*res, tmpRes)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i-1] == nums[i] {
			continue
		}
		roundRes = append(roundRes, nums[i])
		backtrack90(res, nums, roundRes, i+1)
		roundRes = roundRes[:len(roundRes)-1]
	}
}
