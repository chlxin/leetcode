package chapter

import "testing"

func TestIsSameTree(t *testing.T) {
	p := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	q := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	sameTree := isSameTree(p, q)
	if sameTree != true {
		t.Fail()
	}
}


func TestIsSameTree2(t *testing.T) {
	p := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	q := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	sameTree := isSameTree(p, q)
	if sameTree != false {
		t.Fail()
	}
}