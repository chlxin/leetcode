package chapter

import (
	"fmt"
	"testing"
)

func TestAddTwoSum(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	numbers := addTwoNumbers(l1, l2)
	pprint(numbers)
}

func TestAddTwoSum2(t *testing.T) {
	l1 := &ListNode{1, &ListNode{8, nil}}
	l2 := &ListNode{0, nil}
	numbers := addTwoNumbers(l1, l2)
	pprint(numbers)
}

func TestAddTwoSum3(t *testing.T) {
	l1 := &ListNode{5, nil}
	l2 := &ListNode{5, nil}
	numbers := addTwoNumbers(l1, l2)
	pprint(numbers)
}

func pprint(l *ListNode) {
	if l == nil {
		fmt.Println(l)
		return
	}
	for l != nil {
		fmt.Println(*l)
		l = l.Next
	}
}

func TestSS(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	changeA(nums)
	fmt.Println(nums)
}

func changeA(a []int) {
	a[2] = 10
	a = append(a, 20)
}
