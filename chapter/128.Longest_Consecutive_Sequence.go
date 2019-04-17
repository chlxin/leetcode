package chapter

func longestConsecutive(nums []int) int {
	res := 0
	exists := make(map[int]bool)
	for _, v := range nums {
		exists[v] = true
	}

	for _, num := range nums {
		if !exists[num-1] {
			seq, count := num, 0
			for exists[seq] {
				seq++
				count++
			}
			if count > res {
				res = count
			}
		}
	}
	return res
}
