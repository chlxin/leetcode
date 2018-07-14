package chapter

import (
	"testing"
	"fmt"
)

func TestMinDepth(t *testing.T) {
	depth := minDepth(&TreeNode{1, &TreeNode{2, nil, nil}, nil})
	fmt.Println(depth)
}