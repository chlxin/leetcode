package chapter

// use algorithm of greedy
func canJump(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}
	max := 0
	left, right := 0, 0
	for left <= right {
		for i := left; i <= right; i++ {
			if i+nums[i] > max {
				max = i + nums[i]
				if max >= n-1 {
					return true
				}
			}
		}
		left = right + 1
		right = max
	}

	return false
}
