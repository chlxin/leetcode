package chapter

// 用递归做的可能会有栈深度过深导致溢出的风险，其实可以用前序遍历生成一个串来比较串判断两颗树是否相等
// 生成串来比较的方式也有点不好，假如两个串实际上在第二个元素就不一样，其实没有遍历后面的必要
// 所以可以用迭代的先序遍历方式来重写，也就是个先序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameSubTree(p, q)
}

func isSameSubTree(p *TreeNode, q *TreeNode) bool {
	//同时为空相等
	if p == nil && q == nil {
		return true
	}
	//其中有一个为空，那么不相等
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameSubTree(p.Left, q.Left) && isSameSubTree(p.Right, q.Right)
	}
	return false
}
