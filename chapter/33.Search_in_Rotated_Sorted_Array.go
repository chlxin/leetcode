package chapter

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	// lo 就是那个pivot
	pivot := lo
	//执行真正的二分搜索，实际数据可以看成向右移动pivot个元素
	lo, hi = pivot, len(nums)-1+pivot
	for lo < hi {
		mid := (lo + hi) / 2
		if target == nums[mid%len(nums)] {
			return mid % len(nums)
		} else if target < nums[mid%len(nums)] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	if target == nums[lo%len(nums)] {
		return lo % len(nums)
	}
	return -1
}
