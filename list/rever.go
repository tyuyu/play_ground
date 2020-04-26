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

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//
//
//示例：
//
//给你这个链表：1->2->3->4->5
//
//当 k = 2 时，应当返回: 2->1->4->3->5
//
//当 k = 3 时，应当返回: 3->2->1->4->5
//
//
//
//说明：
//
//你的算法只能使用常数的额外空间。
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	np := &ListNode{}
	cur, it := np, head
	for it != nil {
		heap := make([]*ListNode, 0)
		for i := 0; i < k; i++ {
			if it == nil {
				break
			}
			heap = append(heap, it)
			it = it.Next
		}
		if len(heap) == k {
			for j := k - 1; j >= 0; j-- {
				cur.Next = heap[j]
				cur = cur.Next
				if j == 0 {
					cur.Next = nil
				}
			}
		} else {
			cur.Next = heap[0]
		}
	}
	return np.Next
}
