package mysort

import "sort"

//摆动排序 II
//给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
//
//示例 1:
//
//输入: nums = [1, 5, 1, 1, 6, 4]
//输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]
//示例 2:
//
//输入: nums = [1, 3, 2, 2, 3, 1]
//输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
//说明:
//你可以假设所有输入都会得到有效的结果。
//
//进阶:
//你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？

func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	sort.Ints(nums)
	middle := len(nums) / 2
	i, j := 1, middle+1
	for i < middle {
		nums[i], nums[j] = nums[j], nums[i]
		i += 2
		j += 2
	}
}

func findDuplicate(nums []int) int {
	m := make(map[int]struct{}, 0)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = struct{}{}
	}
	return 0
}

func countSmaller(nums []int) []int {
	count := make([]int, len(nums))
	for i, num := range nums {
		if i == len(nums)-1 {
			break
		}
		c := 0
		for _, i3 := range nums[i:] {
			if i3 < num {
				c++
			}
		}
		count[i] = c
	}

	return count
}
