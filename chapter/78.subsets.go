package chapter

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	roundRes := make([]int, 0)
	backtrack(&res, nums, roundRes, 0)
	return res
}

// backtrack 回溯方法，需要track几个变量，总的结果集res，目前迭代的index
// 预定义的nums数组，当前的遍历的结果
// 本质上还是遍历，当进入到backtrack
func backtrack(res *[][]int, nums []int, roundRes []int, index int) {
	tmpRes := make([]int, len(roundRes))
	copy(tmpRes, roundRes)
	*res = append(*res, tmpRes)
	for i := index; i < len(nums); i++ {
		roundRes = append(roundRes, nums[i])
		backtrack(res, nums, roundRes, i+1)
		roundRes = roundRes[:len(roundRes)-1]
	}
}

func subsetsDP(nums []int) [][]int {
	res := dp39(nums)
	return res
}

func dp39(candidates []int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		res = append(res, []int{})
		return res
	}
	// choose c(0)
	choose := dp39(candidates[1:])
	for _, v := range choose {
		copyV := make([]int, len(v))
		copy(copyV, v)
		tmpList := append(copyV, candidates[0])
		res = append(res, tmpList)
	}
	// not choose c(0)
	for _, v := range choose {
		copyV := make([]int, len(v))
		copy(copyV, v)
		res = append(res, copyV)
	}
	return res
}
