package chapter

import (
	"fmt"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	res := searchMatrix(matrix, 11)
	fmt.Println(res)
}
