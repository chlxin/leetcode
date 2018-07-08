package chapter

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		m[n] = i
	}

	for i, n := range nums {
		if s, ok := m[target-n]; ok {
			if i == s {
				continue
			}
			return []int{i, s}
		}
	}
	return []int{}
}
