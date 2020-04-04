package array

import "math/rand"

//打乱一个没有重复元素的数组。
//
//示例:
//
//// 以数字集合 1, 2 和 3 初始化数组。
//int[] nums = {1,2,3};
//Solution solution = new Solution(nums);
//
//// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
//solution.shuffle();
//
//// 重设数组到它的初始状态[1,2,3]。
//solution.reset();
//
//// 随机返回数组[1,2,3]打乱后的结果。
//solution.shuffle();

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	tmp := make([]int, 0)
	tmp = append(tmp, this.nums...)
	rand.Seed(int64(len(this.nums)))
	for i := len(this.nums) - 1; i >= 0; i-- {
		j := rand.Intn(i)
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return tmp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
