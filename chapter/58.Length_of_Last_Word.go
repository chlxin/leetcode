package chapter

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}
	s = s[:i+1]

	cur, llen := 0, 0
	for _, v := range s {
		if v == ' ' {
			cur = 0
			continue
		}
		cur++
		if cur > llen {
			llen = cur
		}
	}
	return cur
}
