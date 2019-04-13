package chapter

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m < 1 {
		return 1
	}
	n := len(obstacleGrid[0])
	if n < 1 {
		return 1
	}
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	// init
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// m-1 row
	for i := n - 1; i >= 0; i-- {
		// 障碍物
		if obstacleGrid[m-1][i] == 1 || (i+1 <= n-1 && dp[m-1][i+1] == 0) {
			dp[m-1][i] = 0
		} else {
			dp[m-1][i] = 1
		}
	}
	// n-1 column
	for i := m - 1; i >= 0; i-- {
		// 障碍物
		if obstacleGrid[i][n-1] == 1 || (i+1 <= m-1 && dp[i+1][n-1] == 0) {
			dp[i][n-1] = 0
		} else {
			dp[i][n-1] = 1
		}
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				donwStep := dp[i+1][j]
				rightStep := dp[i][j+1]
				dp[i][j] = donwStep + rightStep
			}
		}
	}
	return dp[0][0]
}
