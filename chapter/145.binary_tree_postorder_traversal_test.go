package chapter

import (
	"testing"
	"fmt"
)

func TestStack1_PopostorderTraversal(t *testing.T) {
	traversal := postorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}})
	fmt.Println(traversal)

}
