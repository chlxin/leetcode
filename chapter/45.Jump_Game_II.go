package chapter

// dp(i) = min(dp[i+1],...,dp[i+nums[i]]) + 1, exclude dp[k] != 0 and if min(dp[i+1],...,dp[i+nums[i]] > 0
// if min(dp[i+1],...,dp[i+nums[i]]) is 0, then dp[i] = 0
func jump(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	dp := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if nums[i] == 0 {
			dp[i] = 0
		} else if i+nums[i] >= n-1 {
			dp[i] = 1
		} else {
			rightMost := i + nums[i] + 1
			min := min45(dp[i+1 : rightMost])
			if min == 0 {
				dp[i] = 0
			} else {
				dp[i] = min + 1
			}

		}

	}
	return dp[0]
}

func min45(arg []int) int {
	min := 0
	for _, v := range arg {
		if v > 0 {
			if min == 0 {
				min = v
			} else {
				if v < min {
					min = v
				}
			}
		}

	}
	return min
}
