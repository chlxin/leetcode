package chapter

import (
	"fmt"
	"testing"
)

func Test_minCut(t *testing.T) {
	res := minCut("abbab")
	fmt.Println(res)
}
