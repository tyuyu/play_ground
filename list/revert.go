package list

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	pre, i := &ListNode{Next: head}, head
	c := 1
	for c < m {
		c++
		pre = pre.Next
		i = i.Next
	}

	cur := i
	var j *ListNode
	for c <= n && cur.Next != nil {
		c++
		tmp := cur
		cur = cur.Next
		tmp.Next = j
		j = tmp
	}
	i.Next = cur
	pre.Next = j
	return head

}
