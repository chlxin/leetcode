package chapter

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := maxDepth(root.Left), maxDepth(root.Right)
	max := left
	if max < right {
		max = right
	}
	return max + 1
}
