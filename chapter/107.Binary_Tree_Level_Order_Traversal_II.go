package chapter

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var node *TreeNodeWrap
	q := make([]*TreeNodeWrap, 0)
	q = append(q, &TreeNodeWrap{Node: root, Level: 0})
	bucket := make([]*TreeNodeWrap, 0)
	res := make([][]int, 0)
	for len(q) > 0 {
		node, q = q[0], q[1:]
		bucket = append(bucket, node)
		if node.Node.Left != nil {
			q = append(q, &TreeNodeWrap{Node: node.Node.Left, Level: node.Level + 1})
		}
		if node.Node.Right != nil {
			q = append(q, &TreeNodeWrap{Node: node.Node.Right, Level: node.Level + 1})
		}
		for len(res) < node.Level+1 {
			res = append(res, make([]int, 0))
		}
		res[node.Level] = append(res[node.Level], node.Node.Val)
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
