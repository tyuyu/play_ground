package dynamic

import (
	"math"
	"sort"
)

//示例:
//
//int k = 3;
//int[] arr = [4,5,8,2];
//KthLargest kthLargest = new KthLargest(3, arr);
//kthLargest.add(3);   // returns 4
//kthLargest.add(5);   // returns 5
//kthLargest.add(10);  // returns 5
//kthLargest.add(9);   // returns 8
//kthLargest.add(4);   // returns 8

type KthLargest struct {
	k    int
	heap []int
	kv   int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	kl := KthLargest{
		k:    k,
		heap: nums,
	}
	if len(kl.heap) > k {
		kl.heap = kl.heap[len(kl.heap)-k:]
		kl.kv = kl.heap[0]
	} else if len(kl.heap) == k {
		kl.kv = kl.heap[0]
	} else {
		for len(kl.heap) < k {
			kl.heap = append([]int{math.MinInt64}, kl.heap...)
		}
		kl.kv = math.MinInt64
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if val < this.kv {
		return this.kv
	}
	i := find(0, len(this.heap)-1, this.heap, val)
	if i == -1 {
		this.heap = append(this.heap[1:], val)
	} else {
		this.heap = append(append(append([]int{}, this.heap[1:i+1]...), val), this.heap[i+1:]...)
	}
	this.kv = this.heap[0]
	return this.heap[0]
}

func find(bg int, end int, heap []int, val int) int {

	if bg == end {
		return bg
	}

	i := bg + end/2

	if i == len(heap)-1 {
		if val > heap[i] {
			return -1
		} else {
			return i
		}
	}

	if val > heap[i] {
		if val <= heap[i+1] {
			return i
		}
		return find(i+1, end, heap, val)
	} else {
		return find(bg, i-1, heap, val)
	}

}
