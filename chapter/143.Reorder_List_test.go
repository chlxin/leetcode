package chapter

import (
	"fmt"
	"testing"
)

func Test_reorderList(t *testing.T) {
	// 1->2->3->4
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	reorderList(node)
	fmt.Println(node)
}

func Test_reorderList2(t *testing.T) {
	// 1->2->3->4
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	reorderList(node)
	fmt.Println(node)
}
