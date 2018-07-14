package chapter

type Node struct {
	node *TreeNode
	status int
}

type stack1 []*Node

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	res := make([]int, 0)
	var x *Node
	s := make(stack1, 0)
	s = s.Push(&Node{root, 0})
	for !s.isEmpty() {
		s, x = s.Pop()
		switch x.status {
		case 0:
			x.status = 1
			s = s.Push(x)
			if x.node.Left != nil {
				s = s.Push(&Node{x.node.Left, 0})
			}
		case 1:
			x.status = 2
			s = s.Push(x)
			if x.node.Right != nil {
				s = s.Push(&Node{x.node.Right, 0})
			}
		case 2:
			res = append(res, x.node.Val)
		default:
			panic("走到这算我输")
		}
	}
	return res

}


func (s stack1) Push(v *Node) stack1 {
	return append(s, v)
}

func (s stack1) Pop() (stack1, *Node) {
	l := len(s)
	if l == 0 {
		panic("pop an empty stack")
	}
	return  s[:l-1], s[l-1]
}

func (s stack1) isEmpty() bool {
	return len(s) == 0
}
