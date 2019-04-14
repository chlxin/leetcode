package chapter

func numDecodings(s string) int {
	return dynamicProgramming91(s)
}

func recursive91(s string) int {
	if len(s) == 0 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	step1, step2 := 0, 0
	if s[0] == '1' || (s[0] == '2' && s[1] <= '6' && s[1] >= '0') {
		step2 = recursive91(s[2:])
	}
	step1 = recursive91(s[1:])
	return step1 + step2
}

func dynamicProgramming91(s string) int {
	if len(s) == 0 {
		return 1
	}

	dp := make([]int, len(s)+1)
	n := len(s)
	dp[n] = 1
	if s[n-1] != '0' {
		dp[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] == '0' {
			continue // cause dp[i] = 0
		}
		step1, step2 := dp[i+1], 0
		if s[i] == '1' || (s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6') {
			step2 = dp[i+2]
		}
		dp[i] = step1 + step2
	}
	return dp[0]
}
