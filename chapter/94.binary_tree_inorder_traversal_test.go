package chapter

import (
	"testing"
	"fmt"
)

func TestInorderTraversal(t *testing.T)  {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	traversal := inorderTraversal(root)
	fmt.Println(traversal)

}