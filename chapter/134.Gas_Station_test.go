package chapter

import (
	"fmt"
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	res := canCompleteCircuit(gas, cost)
	fmt.Println(res)
}
