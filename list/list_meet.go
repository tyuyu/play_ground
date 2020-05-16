package list

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		if fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}
