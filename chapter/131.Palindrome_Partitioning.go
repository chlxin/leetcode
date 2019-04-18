package chapter

func partition(s string) [][]string {
	bs := []byte(s)
	res := make([][]string, 0)
	backtrack131(&res, make([]string, 0), bs, 0)
	return res
}

func backtrack131(res *[][]string, curList []string, bs []byte, start int) {
	if len(curList) > 0 && start >= len(bs) {
		// last position
		tmpList := make([]string, len(curList))
		copy(tmpList, curList)
		*res = append(*res, tmpList)
	}

	for i := start; i < len(bs); i++ {
		if isPalindrome131(bs, start, i) {
			curList = append(curList, string(bs[start:i+1]))
			backtrack131(res, curList, bs, i+1)
			curList = curList[:len(curList)-1]
		}
	}
}

func isPalindrome131(str []byte, l, r int) bool {
	if l == r {
		return true
	}
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}
