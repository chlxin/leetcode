package chapter

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		res = append(res, make([]int, i+1))
		res[i][0] = 1
		res[i][len(res[i])-1] = 1
	}
	for i := 2; i < numRows; i++ {
		for j := 1; j <= len(res[i])-2; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
