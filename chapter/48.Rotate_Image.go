package chapter

// 如果不考虑in-place的解法，那么转换前后是有关系(i,j) = (j, n-1-j)
func rotate(matrix [][]int) {
	n := len(matrix)
	if n < 2 {
		return
	}
	// reverse by index
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
