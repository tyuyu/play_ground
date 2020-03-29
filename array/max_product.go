package array

import "math"

//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字）。
//
//
//
//示例 1:
//
//输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//示例 2:
//
//输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	//动态规划的核心是就是缓存子集的计算结果呗
	//这里只需要存上一次的结果，可以省不少内存

	max := nums[0]
	min := nums[0]

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		v1 := cur * max
		v2 := cur * min
		max = int(math.Max(math.Max(float64(cur), float64(v1)), float64(v2)))
		min = int(math.Min(math.Min(float64(cur), float64(v1)), float64(v2)))
		if max > res {
			res = max
		}
	}

	return res
}
