package chapter

import "strconv"

// 整体还是递归的思路，对于每一个最高位为i的组合有(n-1)!种，因此对k值计算，可以确定最高位
// 计算中剩余的排序，继续递归
func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return getPremutationRecursive(nums, n, k)
}

func getPremutationRecursive(nums []int, n int, k int) string {
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	perGroupNumber := factorial(n - 1)
	highPos := (k - 1) / perGroupNumber
	num := nums[highPos]
	nums = append(nums[:highPos], nums[highPos+1:]...)
	k = k % perGroupNumber
	if k == 0 {
		k = perGroupNumber
	}
	return strconv.Itoa(num) + getPremutationRecursive(nums, n-1, k)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}
