package array

import "sort"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//
//
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	check := make(map[int]struct{})
	for i := range nums {
		if _, ok := check[nums[i]]; ok {
			continue
		}
		if nums[i] > 0 {
			break
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				now := nums[j]
				j++
				for j < k && nums[j] == now {
					j++
				}
				now = nums[k]
				k--
				for j < k && nums[k] == now {
					k--
				}
			} else if nums[i]+nums[j]+nums[k] < 0 {
				now := nums[j]
				j++
				for j < k && nums[j] == now {
					j++
				}
			} else {
				now := nums[k]
				k--
				for j < k && nums[k] == now {
					k--
				}
			}
		}
		check[nums[i]] = struct{}{}
	}
	return res
}
