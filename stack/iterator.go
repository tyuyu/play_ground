package stack

//
////给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
////
////列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。
////
////
////
////示例 1:
////
////输入: [[1,1],2,[1,1]]
////输出: [1,1,2,1,1]
////解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
////示例 2:
////
////输入: [1,[4,[6]]]
////输出: [1,4,6]
////解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。
//
//// This is the interface that allows for creating nested lists.
//// You should not implement it, or speculate about its implementation
//type NestedInteger struct {
//}
//
//// Return true if this NestedInteger holds a single integer, rather than a nested list.
//func (this NestedInteger) IsInteger() bool {}
//
//// Return the single integer that this NestedInteger holds, if it holds a single integer
//// The result is undefined if this NestedInteger holds a nested list
//// So before calling this method, you should have a check
//func (this NestedInteger) GetInteger() int {}
//
//// Set this NestedInteger to hold a single integer.
//func (n *NestedInteger) SetInteger(value int) {}
//
//// Set this NestedInteger to hold a nested list and adds a nested integer to it.
//func (this *NestedInteger) Add(elem NestedInteger) {}
//
//// Return the nested list that this NestedInteger holds, if it holds a nested list
//// The list length is zero if this NestedInteger holds a single integer
//// You can access NestedInteger's List element directly if you want to modify it
//func (this NestedInteger) GetList() []*NestedInteger { return nil }
//
//type NestedIterator struct {
//	values []int
//}
//
////
////func Constructor(nestedList []*NestedInteger) *NestedIterator {
////	it := &NestedIterator{values: make([]int, 0)}
////	add(it, nestedList)
////	return it
////}
//
//func add(it *NestedIterator, list []*NestedInteger) {
//	for _, nest := range list {
//		if nest.IsInteger() {
//			it.values = append(it.values, nest.GetInteger())
//		} else {
//			add(it, nest.GetList())
//		}
//	}
//}
//
//func (this *NestedIterator) Next() int {
//	if len(this.values) == 0 {
//		return 0
//	}
//	x := this.values[0]
//	this.values = this.values[1:]
//	return x
//}
//
//func (this *NestedIterator) HasNext() bool {
//	return len(this.values) > 0
//}
