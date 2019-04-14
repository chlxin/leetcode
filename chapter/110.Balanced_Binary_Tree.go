package chapter

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right := depth110(root.Left), depth110(root.Right)
	if left > right {
		left, right = right, left
	}
	return (right-left) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth110(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := depth110(root.Left), depth110(root.Right)
	max := left
	if right > max {
		max = right
	}
	return max + 1
}
