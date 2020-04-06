package stack

//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) -- 将元素 x 推入栈中。
//pop() -- 删除栈顶的元素。
//top() -- 获取栈顶元素。
//getMin() -- 检索栈中的最小元素。
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.

type MinStack struct {
	nums []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		nums: make([]int, 0),
		min:  0,
	}
}

func (this *MinStack) Push(x int) {
	if len(this.nums) == 0 {
		this.min = x
	} else if this.min > x {
		this.min = x
	}
	this.nums = append(this.nums, x)
}

func (this *MinStack) Pop() {
	out := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	if out == this.min {
		if len(this.nums) == 0 {
			this.min = 0
			return
		}
		this.min = this.nums[0]
		for _, num := range this.nums {
			if num < this.min {
				this.min = num
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
