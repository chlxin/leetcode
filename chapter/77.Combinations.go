package chapter

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	backtrack77(&res, make([]int, 0), 1, n, k)
	return res
}

func backtrack77(res *[][]int, nums []int, start, n, k int) {
	if k == 0 {
		tmpRes := make([]int, len(nums))
		copy(tmpRes, nums)
		*res = append(*res, tmpRes)
	}
	for i := start; i <= n; i++ {
		nums = append(nums, i)
		backtrack77(res, nums, i+1, n, k-1)
		nums = nums[:len(nums)-1]
	}
}
