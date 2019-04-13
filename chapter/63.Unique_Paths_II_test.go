package chapter

import (
	"fmt"
	"testing"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	obstacles := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	res := uniquePathsWithObstacles(obstacles)
	fmt.Println(res)
}

func Test_uniquePathsWithObstacles2(t *testing.T) {
	obstacles := [][]int{
		{0, 0},
	}
	res := uniquePathsWithObstacles(obstacles)
	fmt.Println(res)
}

func Test_uniquePathsWithObstacles3(t *testing.T) {
	obstacles := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}
	res := uniquePathsWithObstacles(obstacles)
	fmt.Println(res)
}
