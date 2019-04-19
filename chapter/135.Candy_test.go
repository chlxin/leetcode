package chapter

import (
	"fmt"
	"testing"
)

func Test_candy(t *testing.T) {
	candies := []int{1, 0, 2}
	res := candy(candies)
	fmt.Println(res)
}
