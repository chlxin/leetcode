package chapter

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	roundRes := make([]int, 0)
	backtrack39(&res, candidates, roundRes, 0, target)
	return res
}

func backtrack39(res *[][]int, candidates []int, roundRes []int, index int, remain int) {
	if remain < 0 {
		return
	}
	if remain == 0 {
		tmpRes := make([]int, len(roundRes))
		copy(tmpRes, roundRes)
		*res = append(*res, tmpRes)
		return
	}
	// remain > 0
	for i := index; i < len(candidates); i++ {
		roundRes = append(roundRes, candidates[i])
		backtrack39(res, candidates, roundRes, i, remain-candidates[i])
		roundRes = roundRes[:len(roundRes)-1]
	}
}
