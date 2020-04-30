package m

import (
	"fmt"
	"sort"
)

type LRUCache struct {
	heap []int
	set  map[int]int
	cap  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		heap: make([]int, 0),
		set:  make(map[int]int, 0),
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.set[key]; ok {
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.set[key]; !ok {
		this.heap = append(this.heap, key)
	}
	this.set[key] = value
	sort.Slice(this.heap, func(i, j int) bool {
		return this.set[this.heap[i]] < this.set[this.heap[j]]
	})
	if len(this.heap) > this.cap {
		p := this.heap[0]
		this.heap = this.heap[1:]
		delete(this.set, p)
	}
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	prevStr := "1"
	curChar := "1"
	curCount := 0
	curStr := ""
	for i := 2; i <= n; i++ {
		for _, v := range prevStr {
			vChar := string(v)
			if curChar == vChar {
				curCount++
			} else {
				curStr += fmt.Sprintf("%d%s", curCount, curChar)
				curChar = vChar
				curCount = 1
			}
		}
		prevStr = curStr + fmt.Sprintf("%d%s", curCount, curChar)
		curStr = ""
		curCount = 0
		curChar = string(prevStr[0])
	}
	return prevStr
}
