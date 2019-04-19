package chapter

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.Next
		if fast.Next == nil || fast.Next.Next == nil {
			return nil // 没有循环
		}
		fast = fast.Next.Next
		// 相遇在k点
		if slow == fast {
			break
		}
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
