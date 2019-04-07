package chapter

import (
	"fmt"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	res := romanToInt("MCMXCIV")
	fmt.Println(res)
}
