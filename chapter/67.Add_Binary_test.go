package chapter

import (
	"fmt"
	"testing"
)

func Test_addBinary(t *testing.T) {
	res := addBinary("1", "111")
	fmt.Println(res)
}
