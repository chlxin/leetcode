package chapter

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	roundRes := make([]int, 0)
	backtrack40(&res, candidates, roundRes, 0, target)
	return res
}

func backtrack40(res *[][]int, candidates []int, roundRes []int, index int, remain int) {
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
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		roundRes = append(roundRes, candidates[i])
		backtrack40(res, candidates, roundRes, i+1, remain-candidates[i])
		roundRes = roundRes[:len(roundRes)-1]
	}

}
