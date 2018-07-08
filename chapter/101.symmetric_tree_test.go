package chapter

import "testing"

func TestIsSymmetric(t *testing.T) {
	r := &TreeNode{1,
	&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
	&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}
	res := isSymmetric(r)
	if !res {
		t.Fail()
	}
}
