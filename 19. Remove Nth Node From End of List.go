/*
19. Remove Nth Node From End of List

Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}
//one-pass approach
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	prenode := node
	isHead := false
	for i := 0; i < n; i++ {
		if node.Next != nil {
			node = node.Next
		} else {
			isHead = true
		}
	}
	if isHead {
		return head.Next
	}

	for node.Next != nil {
		node = node.Next
		prenode = prenode.Next
	}
	prenode.Next = prenode.Next.Next

	return head
}