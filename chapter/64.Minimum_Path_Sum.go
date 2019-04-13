package chapter

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	// so m >= 1, n >= 1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = grid[i][j]
			} else if i == m-1 {
				dp[i][j] = dp[i][j+1] + grid[i][j]
			} else if j == n-1 {
				dp[i][j] = dp[i+1][j] + grid[i][j]
			} else {
				min := dp[i+1][j]
				if dp[i][j+1] < min {
					min = dp[i][j+1]
				}
				dp[i][j] = grid[i][j] + min
			}
		}
	}
	return dp[0][0]
}
