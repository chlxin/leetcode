package chapter


func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		 return make([]int, 0)
	}
	res := make([]int, 0)
	var x *TreeNode
	s := make(stack, 0)
	s = s.Push(root)
	for !s.isEmpty() {
		s, x = s.Pop()
		res = append(res, x.Val)
		if x.Right != nil {
			s = s.Push(x.Right)
		}
		if x.Left != nil {
			s = s.Push(x.Left)
		}
	}
	return res
}


func preorderTraversalrecursion(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	res := []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

