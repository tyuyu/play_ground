package array

import (
	"sort"
)

//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
//注意：
//
//答案中不可以包含重复的四元组。
//
//示例：
//
//给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
//满足要求的四元组集合为：
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/4sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	if len(nums) < 4 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, len(nums)-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					for nums[k] == nums[k-1] && k < l {
						k++
					}
					l--
					for nums[l] == nums[l+1] && k < l {
						l--
					}
				}
				if sum > target {
					l--
					for nums[l] == nums[l+1] && k < l {
						l--
					}
				}
				if sum < target {
					k++
					for nums[k] == nums[k-1] && k < l {
						k++
					}
				}
			}
		}
	}

	return res
}
