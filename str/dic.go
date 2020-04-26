package str

import "sort"

//实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须原地修改，只允许使用额外常数空间。
//
//以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
//1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1
//

func nextPermutation(nums []int) {

	if len(nums) < 2 {
		return
	}

	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i+1] > nums[i] {
			break
		}
	}
	if i >= 0 {
		j := len(nums) - 1
		for ; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	sort.Ints(nums[i+1:])
}
