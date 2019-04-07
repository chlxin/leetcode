package chapter

func plusOne(digits []int) []int {
	reverse(digits)
	n := len(digits)
	for i := 0; i < n; i++ {
		digits[i] = digits[i] + 1
		flag := true
		if digits[i] >= 10 {
			flag = false
			digits[i] = digits[i] % 10
			if i == n-1 {
				digits = append(digits, 1)
			}
		}
		if flag {
			break
		}
	}
	reverse(digits)
	return digits
}

// func reverse(nums []int) {
// 	for i := 0; i < len(nums)/2; i++ {
// 		j := len(nums) - 1 - i
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// }
