package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addDigit(l1, l2, 0)
}

func addTwoNumbersImproved(l1 *ListNode, l2 *ListNode) *ListNode {
	var surplus int
	var head, current *ListNode

	for l1 != nil || l2 != nil || surplus != 0 {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + surplus
		surplus = surplus / 10
		if current == nil {
			head = &ListNode{Val: sum % 10, Next: nil}
			current = head
		} else {
			current.Next = &ListNode{Val: sum % 10, Next: nil}
			current = current.Next
		}
	}
	return head
}

func addDigit(l1 *ListNode, l2 *ListNode, add int) *ListNode {
	if l1 == nil && l2 == nil {
		//empty
		if add == 0 {
			return nil
		} else {
			return &ListNode{Val: add, Next: nil}
		}
	} else if l1 == nil {
		sum := l2.Val + add
		if sum >= 10 {
			nextNode := addDigit(nil, l2.Next, 1)
			return &ListNode{Val: sum - 10, Next: nextNode}
		} else {
			nextNode := addDigit(nil, l2.Next, 0)
			return &ListNode{Val: sum, Next: nextNode}
		}
	} else if l2 == nil {
		sum := l1.Val + add
		if sum >= 10 {
			nextNode := addDigit(l1.Next, nil, 1)
			return &ListNode{Val: sum - 10, Next: nextNode}
		} else {
			nextNode := addDigit(l1.Next, nil, 0)
			return &ListNode{Val: sum, Next: nextNode}
		}
	} else {
		sum := l1.Val + l2.Val + add
		if sum >= 10 {
			nextNode := addDigit(l1.Next, l2.Next, 1)
			return &ListNode{Val: sum - 10, Next: nextNode}
		} else {
			nextNode := addDigit(l1.Next, l2.Next, 0)
			return &ListNode{Val: sum, Next: nextNode}
		}
	}
}
