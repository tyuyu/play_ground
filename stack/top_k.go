package stack

import "container/heap"

//给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
//
//示例 1:
//
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//示例 2:
//
//输入: nums = [1], k = 1
//输出: [1]
//说明：
//
//你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
//你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num]++
	}
	h := &freqQueue{}
	for k, v := range m {
		h.Push(freq{
			val:   k,
			count: v,
		})
	}
	heap.Init(h)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).(freq).val
	}
	return res
}

type freq struct {
	val   int
	count int
}

type freqQueue []freq

func (f freqQueue) Len() int {
	return len(f)
}

func (f freqQueue) Less(i, j int) bool {
	return f[i].count > f[j].count
}

func (f freqQueue) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *freqQueue) Push(x interface{}) {
	*f = append(*f, x.(freq))
}

func (f *freqQueue) Pop() interface{} {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}
