package chapter

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

//  通过快慢指针能迅速定位定中位节点，同时记录慢指针的上一节点，然后递归生成树节点
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	var (
		left  *TreeNode
		right *TreeNode
	)
	if prev != nil {
		left = sortedListToBST(head)
	}
	if slow.Next != nil {
		right = sortedListToBST(slow.Next)
	}

	return &TreeNode{
		Val:   slow.Val,
		Left:  left,
		Right: right,
	}
}
