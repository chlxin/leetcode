package chapter

import (
	"testing"
	"fmt"
)

func TestBuildTree2(t *testing.T) {
	tree := buildTree_2([]int{2, 3, 1}, []int{3, 2, 1})
	fmt.Println(tree)
}
