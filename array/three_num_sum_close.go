package array

import (
	"math"
	"sort"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
//例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
//
//与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum-closest
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	min, out := math.MaxInt64, 0
	for i := 0; i < len(nums)-2; i++ {
		if i > 1 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			d := int(math.Abs(float64(sum - target)))
			if d < min {
				min = d
				out = sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return out
}
