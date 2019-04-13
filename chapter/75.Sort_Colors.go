package chapter

func sortColors(nums []int) {
	p1, p2 := 0, len(nums)-1
	index := 0
	for index <= p2 {
		if nums[index] == 0 {
			nums[index] = nums[p1]
			nums[p1] = 0
			p1++
		}
		if nums[index] == 2 {
			nums[index] = nums[p2]
			nums[p2] = 2
			p2--
			index--
		}
		index++
	}
}
