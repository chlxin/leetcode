package chapter

// dp[i][j]记录在(i,j)点开始到(m-1, n-1)唯一的走法
// 有递推关系: dp[i][j]=dp[i+1][j]+dp[i][j+1]  (if i+1 < m, j+1 < n)
// dp[m-1][n-1]=0
// 有更简单的方法，实际上m-1行和n-1列都是1，可以预赋值，之前没想到，改也很容易，懒得改了，递推公式是对的都好说
func uniquePaths(m int, n int) int {
	if m < 2 && n < 2 {
		return 1
	}
	// init
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// dp[m-1][n-2], dp[m-2][n-1] = 1, 1 // m, n > 2
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-2 {
				dp[i][j] = 1
				continue
			}
			if i == m-2 && j == n-1 {
				dp[i][j] = 1
				continue
			}
			donwStep, rightStep := 0, 0
			if i+1 < m {
				donwStep = dp[i+1][j]
			}
			if j+1 < n {
				rightStep = dp[i][j+1]
			}
			dp[i][j] = donwStep + rightStep
		}
	}
	return dp[0][0]
}
