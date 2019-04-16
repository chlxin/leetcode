package chapter

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	total := sumNaturalNumber(m)
	res := make([]int, total)
	// last row
	for j := 0; j <= m-1; j++ {
		res[helpIndex(m-1, j)] = triangle[m-1][j]
	}
	for i := m - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			min := res[helpIndex(i+1, j)]
			next := res[helpIndex(i+1, j+1)]
			if min > next {
				min = next
			}
			res[helpIndex(i, j)] = min + triangle[i][j]
		}
	}
	return res[0]
}

func helpIndex(i, j int) int {
	ll := sumNaturalNumber(i)
	return ll + j
}

func sumNaturalNumber(i int) int {
	return (i * (i + 1)) / 2
}
