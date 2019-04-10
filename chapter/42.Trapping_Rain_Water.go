package chapter

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	res := 0
	for left < right {
		// 确保了右边一定有比左边打的数存在，或者相反，很重要！
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}
	return res
}
