package chapter

func preorderTraversal144(root *TreeNode) []int {
	res := make([]int, 0)
	cur := root
	s := make(stack, 0)
	for cur != nil {
		s = s.Push(cur)
		cur = cur.Left
	}

	for !s.isEmpty() {
		s, cur = s.Pop()
		res = append(res, cur.Val)
		// 处理右元素
		if cur.Right != nil {
			cur = cur.Right
			for cur != nil {
				s = s.Push(cur)
				cur = cur.Left
			}
		}
	}
	return res
}
