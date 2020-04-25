package list

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	cur := head
	last := cur
	for i := 0; i < n; i++ {
		last = last.Next
	}

	if last == nil {
		//delete head
		return head.Next
	}

	for last.Next != nil {
		last = last.Next
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return head

}
