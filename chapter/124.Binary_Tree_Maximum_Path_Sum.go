package chapter

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MinInt64
	dfs124(root, &res)
	return res
}

// 该dfs返回的是包含node的单边最长和. 在搜索的过程中需要同时更新maxSum
// 需要比较传入的maxSum,与左右子树+node节点之和比较，更新maxSum
func dfs124(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	left := max(0, dfs124(node.Left, maxSum))
	right := max(0, dfs124(node.Right, maxSum))
	*maxSum = max(*maxSum, node.Val+left+right)
	return max(left, right) + node.Val
}
