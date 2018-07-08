package chapter

//给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
//
//你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

/**
* Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{0, nil}
	cur := res
	//x := reverse(l1)
	//y := reverse(l2)
	flag := 0
	for flag == 1 || l1 != nil || l2 != nil {
		var a, b int
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		s := a + b + flag
		if s >= 10 {
			flag = 1
			s = s - 10
		} else {
			flag = 0
		}
		cur.Next = &ListNode{s, nil}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return res.Next
}

//刚开始看题意，以为是正序的，所以写了个逆序的方法
//func reverse(l *ListNode) *ListNode {
//	if l == nil || l.Next == nil {
//		return l
//	}
//	var prev, head *ListNode = nil, l
//	for head != nil {
//		next := head.Next
//		head.Next = prev
//		prev, head = head, next
//	}
//	return prev
//}
