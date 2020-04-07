package stack

import (
	"container/heap"
)

//给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第k小的元素。
//请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
//
//
//
//示例:
//
//matrix = [
//   [ 1,  5,  9],
//   [10, 11, 13],
//   [12, 13, 15]
//],
//k = 8,
//
//返回 13。
//
//
//提示：
//你可以假设 k 的值永远是有效的, 1 ≤ k ≤ n2 。

func kthSmallest(matrix [][]int, k int) int {
	//最大堆
	h := &IntHeap{}
	for _, ints := range matrix {
		for _, i := range ints {
			h.Push(i)
		}
	}
	heap.Init(h)
	count := 0
	for h.Len() > 0 {
		v := heap.Pop(h).(int)
		count++
		if count == k {
			return v
		}
	}
	return 0
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
