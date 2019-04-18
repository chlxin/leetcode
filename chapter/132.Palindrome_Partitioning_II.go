package chapter

// copy
// cut[i] is the minimum of cut[j - 1] + 1 (j <= i), if [j, i] is palindrome.
func minCut(s string) int {
	bs := []byte(s)
	n := len(bs)
	cut := make([]int, n)
	pals := make([][]bool, n)
	for i := range pals {
		pals[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		// 最小的情况就是前i+1个只需要i刀
		min := i
		for j := 0; j <= i; j++ {
			// determine (j,i),(both include) is Palindrome
			// bs[j] == bs[i] && pals[j+1][i-1] 实际上是一次优化的剪枝，可以避免很多次的额外搜索
			if bs[j] == bs[i] && (j+1 > i-1 || pals[j+1][i-1]) {
				pals[j][i] = true
				if j == 0 {
					min = 0
				} else {
					if cut[j-1]+1 < min {
						min = cut[j-1] + 1
					}
				}
			}
		}
		cut[i] = min
	}
	return cut[n-1]
}
