package chapter

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sum := nums[0]
	max := sum
	n := len(nums)
	for i := 1; i < n; i++ {
		if sum < 0 {
			sum = 0
		}
		sum = sum + nums[i]
		if sum > max {
			max = sum
		}
	}
	return max
}
