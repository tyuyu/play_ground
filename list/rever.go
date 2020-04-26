package list

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3.

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	np := &ListNode{}
	cur := np
	it1, it2 := head, head.Next
	for it1 != nil {
		it2 = it1.Next
		if it2 == nil {
			cur.Next = it1
			it1.Next = nil
			break
		}
		tmp1, tmp2 := it1, it2
		it1 = it2.Next
		cur.Next = tmp2
		tmp2.Next = tmp1
		cur = tmp1
		tmp1.Next = nil
	}
	return np.Next
}
