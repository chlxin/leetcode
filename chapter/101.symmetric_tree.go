package chapter


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirrorIterate(root.Left, root.Right)
}

// 递归版本
func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// 非递归版本
func isMirrorIterate(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	var x, y *TreeNode
	q1 := make([]*TreeNode, 0)
	q2 := make([]*TreeNode, 0)
	q1 = append(q1, left)
	q2 = append(q2, right)
	//进行迭代，每次都从q1和q2中取出对应的元素比较，假如不等直接gg，假如相等，看左子树和右子树，需要比较部分放入两个队列
	for len(q1) > 0 && len(q2) > 0 {
		x, q1 = q1[0], q1[1:]
		y, q2 = q2[0], q2[1:]
		if x.Val != y.Val {
			return false
		}
		if (x.Left != nil) != (y.Right != nil) {
			return false
		}
		if (x.Right != nil) != (y.Left != nil) {
			return false
		}
		if x.Left != nil && y.Right != nil {
			q1 = append(q1, x.Left)
			q2 = append(q2, y.Right)
		}
		if x.Right != nil && y.Left != nil {
			q1 = append(q1, x.Right)
			q2 = append(q2, y.Left)
		}
	}
	return true;
}
