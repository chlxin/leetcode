package chapter

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var x *TreeNode
	res := 0
	s := make(stack, 0)
	s = s.Push(root)
	for !s.isEmpty() {
		s, x = s.Pop()
		if x.Right != nil {
			s = s.Push(x.Right)
		}
		if x.Left != nil {
			if x.Left.Left == nil && x.Left.Right == nil {
				res += x.Left.Val
			}
			s = s.Push(x.Left)
		}
	}
	return res

}

