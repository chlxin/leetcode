package chapter

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	halfHead := slow.Next
	slow.Next = nil
	head1 := head
	head2 := reverseList143(halfHead)
	var prev *ListNode
	for head1 != nil || head2 != nil {
		if head1 != nil && head2 != nil {
			next1 := head1.Next
			next2 := head2.Next

			head1.Next = head2
			head2.Next = next1
			prev = head2

			head1 = next1
			head2 = next2
		} else if head2 != nil {
			prev.Next = head2
			head2 = head2.Next
		} else if head1 != nil {
			prev.Next = head1
			head1 = head1.Next
		}
	}

}

func reverseList143(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	var prev, head *ListNode = nil, l
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}
