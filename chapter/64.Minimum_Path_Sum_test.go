package chapter

import (
	"fmt"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	res := minPathSum(grid)
	fmt.Println(res)
}

func Test_minPathSum2(t *testing.T) {
	grid := [][]int{
		{1, 2, 5},
		{3, 2, 1},
	}
	res := minPathSum(grid)
	fmt.Println(res)
}
