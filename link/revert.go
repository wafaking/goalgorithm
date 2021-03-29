package link

// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur = head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
