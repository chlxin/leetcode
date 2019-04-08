package chapter

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// find 1st a[i]>target
	low, high := -1, -1
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		left++
	}
	high = left
	// find 1st a[i]<target
	left, right = 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] == target {
		left--
	}
	low = left
	if high-low <= 1 {
		return []int{-1, -1}
	}
	return []int{low + 1, high - 1}
}
