package chapter

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				x := 2
				if i-2 > 0 {
					x = dp[i-2] + 2
				}
				dp[i] = x
				if dp[i] > res {
					res = dp[i]
				}
			} else if s[i-1] == ')' {
				l := i - dp[i-1] - 1
				if l >= 0 && s[l] == '(' {
					x := 0
					if l-1 >= 0 {
						x = dp[l-1]
					}
					dp[i] = dp[i-1] + 2 + x
					if dp[i] > res {
						res = dp[i]
					}
				}
			} else {
				panic("illegal input")
			}
		}
	}
	return res
}
