package chapter

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}
	// 保证rating中i比i-1要大的，candy一定i比i-1要大
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 保证rating中i比i+1要大的，candy一定i比i+1要大,特别注意，candy只增不减，就能保证前边的性质依然成立
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	res := 0
	for _, candy := range candies {
		res += candy
	}
	return res
}
