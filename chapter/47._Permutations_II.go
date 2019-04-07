package chapter

import "sort"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	roundRes := make([]int, 0)
	notUse := make(map[int]bool)
	for i := range nums {
		notUse[i] = true
	}
	sort.Ints(nums)
	backtrack47(&res, nums, roundRes, notUse)
	return res
}

// notUse 实际上把notUse当做set来使用，如果未使用就是true，已经使用就是false
func backtrack47(res *[][]int, nums []int, roundRes []int, notUse map[int]bool) {
	if len(notUse) == 0 {
		tmpRes := make([]int, len(roundRes))
		copy(tmpRes, roundRes)
		*res = append(*res, tmpRes)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !notUse[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && notUse[i-1] {
			continue
		}
		roundRes = append(roundRes, nums[i])
		// mark index=i use
		delete(notUse, i)
		backtrack47(res, nums, roundRes, notUse)
		// recover index=i not use
		roundRes = roundRes[:len(roundRes)-1]
		notUse[i] = true
	}
}
