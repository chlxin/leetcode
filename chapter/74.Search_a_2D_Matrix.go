package chapter

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	// 先在col=0列做二分查找，定位在哪一行后，继续用二分查找 (这里默认向下增长了，可能起名字有点反直觉)
	top, bottom := 0, m-1
	for top < bottom {
		mid := (top + bottom) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			top = mid + 1
		} else {
			bottom = mid
		}
	}
	cur := top
	if cur-1 >= 0 && matrix[cur][0] > target {
		cur--
	}
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if matrix[cur][mid] == target {
			return true
		} else if matrix[cur][mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if matrix[cur][left] == target {
		return true
	}
	return false
}
