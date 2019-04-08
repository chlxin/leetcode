package chapter

import "sort"

// 主要思路还是先选择好两个点，然后在剩下的数组里寻找两个值，使得4个值之和为target
// 注意去重，以及适当的剪枝
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i <= n-4; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j <= n-3; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			t2 := target - nums[i] - nums[j]
			for left < right {
				if nums[left]+nums[right] == t2 {
					oneOfRes := []int{nums[i], nums[j], nums[left], nums[right]}
					res = append(res, oneOfRes)
					// 过滤掉重复的元素
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[left]+nums[right] < t2 {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
