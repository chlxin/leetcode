package chapter

type middleResult struct {
	level int
	node *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	var (
		cur *middleResult
		res = make([][]int, 0)
	)

	q := make([]*middleResult, 0)
	q = append(q, &middleResult{0, root})
	for len(q) > 0 {
		cur, q = q[0], q[1:]
		// 其实这段用到了层数是一层一层遍历的trick，写法本身是有些不好的地方，可以改善
		if cur.level >= len(res){
			res = append(res, make([]int, 0))
		}
		res[cur.level] = append(res[cur.level], cur.node.Val)
		if cur.node.Left != nil {
			q = append(q, &middleResult{cur.level + 1, cur.node.Left})
		}
		if cur.node.Right != nil {
			q = append(q, &middleResult{cur.level + 1, cur.node.Right})
		}
	}
	return res
}