package list

import "sort"

//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
//示例:
//
//输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func mergeKLists(lists []*ListNode) *ListNode {
	sr := &ListNode{}
	cur := sr
	nexts := make([]*ListNode, 0)
	for _, node := range lists {
		if node == nil {
			continue
		}
		nexts = append(nexts, node)
	}
	sort.Slice(nexts, func(i, j int) bool {
		return nexts[i].Val < nexts[j].Val
	})
	for len(nexts) > 0 {
		cur.Next = &ListNode{
			Val:  nexts[0].Val,
			Next: nil,
		}
		cur = cur.Next
		tmp := nexts[0]
		nexts = nexts[1:]
		nexts = add(nexts, tmp.Next)
	}
	return sr.Next
}

func add(nexts []*ListNode, next *ListNode) []*ListNode {
	if next == nil {
		return nexts
	}
	index := 0
	for ; index < len(nexts); index++ {
		if nexts[index].Val > next.Val {
			break
		}
	}
	if index == 0 {
		return append([]*ListNode{next}, nexts...)
	}
	if index == len(nexts) {
		return append(nexts, next)
	}
	tmp := make([]*ListNode, 0)
	tmp = append(tmp, nexts[0:index]...)
	tmp = append(tmp, next)
	tmp = append(tmp, nexts[index:]...)
	return tmp
}
