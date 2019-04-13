package chapter

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}
	m, n := len(matrix), len(matrix[0])
	top, right, down, left := len(matrix), len(matrix[0]), 0, -1
	// ->:0, top: 1, <-:2, down:3
	direction := 0
	row, col := 0, 0
	total := m * n
Outer:
	for i := 0; i < total; i++ {
		res = append(res, matrix[row][col])
		// determine row, col
		switch direction {
		case 0:
			if col+1 == right && row+1 == top {
				break Outer
			}
			if col+1 == right {
				direction = 1
				row++
				right = col
			} else {
				col++
			}
		case 1:
			if row+1 == top && col-1 == left {
				break Outer
			}
			if row+1 == top {
				direction = 2
				col--
				top = row
			} else {
				row++
			}
		case 2:
			if col-1 == left && row-1 == down {
				break Outer
			}
			if col-1 == left {
				direction = 3
				row--
				left = col
			} else {
				col--
			}
		case 3:
			if row-1 == down && col+1 == right {
				break Outer
			}
			if row-1 == down {
				direction = 0
				col++
				down = row
			} else {
				row--
			}
		}
	}
	return res
}
