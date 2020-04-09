package list

//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

func isPalindrome(head *ListNode) bool {

	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	rs := reverse(slow)
	it1, it2 := rs, head
	for it1 != nil {
		if it1.Val != it2.Val {
			return false
		}
		it1 = it1.Next
		it2 = it2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
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
