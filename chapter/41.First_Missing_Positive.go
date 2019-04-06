package chapter

// 这道题比较trick一些，因为默认所有的数其实是从1开始增长的，那么只要把第i个元素对应的值放到第nums[i]的位置上
// 然后一次遍历，就可以找出那个丢失的数
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
