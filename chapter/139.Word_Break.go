package chapter

// dp[i]代表从0-i(exclusive)中能够在在wordDict中找到元素组合(true/false)
// 因此dp[j]为true的条件为: 存在i, i<j, dp[i]为true，且(i, j)(exclusive)在wordDict
func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, word := range wordDict {
		words[word] = true
	}
	raw := []byte(s)
	dp := make([]bool, len(s)+1)
	dp[0] = true // base case
	for i := 1; i <= len(raw); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && words[string(raw[j:i])] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(raw)]
}
