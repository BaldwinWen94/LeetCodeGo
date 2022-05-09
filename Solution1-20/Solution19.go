package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = &ListNode{Next: head}

	fast, slow := dummy, dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
