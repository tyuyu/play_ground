package list

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var last *ListNode
	for cur.Next != nil {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	cur.Next = last
	return cur
}
