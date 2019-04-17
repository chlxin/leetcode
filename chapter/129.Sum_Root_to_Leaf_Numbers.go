package chapter

func sumNumbers(root *TreeNode) int {
	var total int
	helper(root, 0, &total)
	return total
}

func helper(root *TreeNode, sum int, total *int) {
	if root == nil {
		return
	}

	sum = sum*10 + root.Val
	// leaf
	if root.Left == nil && root.Right == nil {
		*total += sum
		return
	}

	helper(root.Left, sum, total)
	helper(root.Right, sum, total)
}
