package chapter

//思路，先往左路延伸，途中经过的都入栈。
// 之后开始迭代，循环条件是栈中存在活动的元素就要进行；循环不变式，每一次从栈中取出的元素都已经遍历过左子树，
// 因此只要对当前元素先做操作，再处理当前元素右子树，右子树一路遍历到最左，期间元素都入栈

type stack []*TreeNode

func inorderTraversal(root *TreeNode) []int {
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

func (s stack) Push(v *TreeNode) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, *TreeNode) {
	l := len(s)
	if l == 0 {
		panic("pop an empty stack")
	}
	return  s[:l-1], s[l-1]
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}

