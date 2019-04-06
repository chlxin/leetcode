package chapter

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	index := n - 1
	for ; index > 0; index-- {
		if nums[index] > nums[index-1] {
			break
		}
	}
	if index == 0 {
		reverse(nums)
		return
	}
	// index > 0
	for i := n - 1; i >= index; i-- {
		if nums[i] > nums[index-1] {
			nums[i], nums[index-1] = nums[index-1], nums[i]
			break
		}
	}
	reverse(nums[index:])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - 1 - i
		nums[i], nums[j] = nums[j], nums[i]
	}
}
