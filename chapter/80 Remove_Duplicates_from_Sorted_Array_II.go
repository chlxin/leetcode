package chapter

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	probe := 1
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[probe-1] {
			probe++
			nums[probe] = nums[i]
		}
	}
	return probe + 1
}
