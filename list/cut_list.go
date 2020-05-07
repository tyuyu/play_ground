package list

func partition(head *ListNode, x int) *ListNode {

	if head == nil {
		return head
	}
	l, r := &ListNode{}, &ListNode{}
	it1, it2 := l, r
	it := head
	for it != nil {
		tmp := it
		it = it.Next
		if tmp.Val < x {
			it1.Next = tmp
			it1 = it1.Next
			it1.Next = nil
		} else {
			it2.Next = tmp
			it2 = it2.Next
			it2.Next = nil
		}

	}
	it1.Next = r.Next
	return l.Next
}
