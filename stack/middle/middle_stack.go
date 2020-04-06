package middle

//中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
//
//例如，
//
//[2,3,4] 的中位数是 3
//
//[2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
//设计一个支持以下两种操作的数据结构：
//
//void addNum(int num) - 从数据流中添加一个整数到数据结构中。
//double findMedian() - 返回目前所有元素的中位数。
//示例：
//
//addNum(1)
//addNum(2)
//findMedian() -> 1.5
//addNum(3)
//findMedian() -> 2
//进阶:
//
//如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
//如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？

type MedianFinder struct {
	lefts    []int
	rights   []int
	leftMax  int
	rightMin int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		lefts:    make([]int, 0),
		rights:   make([]int, 0),
		leftMax:  0,
		rightMin: 0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.lefts) == 0 {
		this.lefts = append(this.lefts, num)
		this.leftMax = num
		return
	}
	if num < this.leftMax {
		this.lefts = insert(this.lefts, num)
		if len(this.lefts)-1 > len(this.rights) {
			this.rights = append([]int{this.lefts[len(this.lefts)-1]}, this.rights...)
			this.lefts = this.lefts[:len(this.lefts)-1]
		}
	} else {
		this.rights = insert(this.rights, num)
		if len(this.rights) > len(this.lefts) {
			this.lefts = append(this.lefts, this.rights[0])
			this.rights = this.rights[1:]
		}
	}
	this.leftMax = this.lefts[len(this.lefts)-1]
	this.rightMin = this.rights[0]
}

func insert(nums []int, num int) []int {
	index := 0
	for ; index < len(nums); index++ {
		if num < nums[index] {
			break
		}
	}
	left := nums[:index]
	right := append([]int{num}, nums[index:]...)
	return append(left, right...)
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.lefts) > len(this.rights) {
		return float64(this.leftMax)
	} else {
		return float64(this.leftMax+this.rightMin) / 2
	}
}
