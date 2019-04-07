package chapter

import (
	"fmt"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	res := uniquePaths(7, 3)
	fmt.Println(res)
}
