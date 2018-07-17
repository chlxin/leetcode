package chapter

import (
	"testing"
	"fmt"
)

func TestStack1_PopostorderTraversal(t *testing.T) {
	traversal := postorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}})
	fmt.Println(traversal)

}

func TestSumOfLeftLeaves(t *testing.T)  {
	leaves := sumOfLeftLeaves(&TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}})
	fmt.Println(leaves)
}