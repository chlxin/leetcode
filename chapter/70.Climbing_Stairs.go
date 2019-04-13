package chapter

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	dp := make([]int, n)
	dp[n-1], dp[n-2] = 1, 2
	for i := n - 3; i >= 0; i-- {
		dp[i] = dp[i+1] + dp[i+2]
	}
	return dp[0]
}
