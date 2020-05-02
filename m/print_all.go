package m

import "sort"

//给定一个可包含重复数字的序列，返回所有不重复的全排列。
//
//示例:
//
//输入: [1,1,2]
//输出:
//[
//  [1,1,2],
//  [1,2,1],
//  [2,1,1]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutations-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func permuteUnique(nums []int) [][]int {

	res := new([][]int)

	sort.Ints(nums)

	dfs(nums, []int{}, res)

	return *res
}

func dfs(nums []int, old []int, res *[][]int) {

	if len(nums) == 1 {
		*res = append(*res, append(old, nums[0]))
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		cur := append(old, nums[i])
		next := make([]int, 0)
		next = append(next, nums[:i]...)
		next = append(next, nums[i+1:]...)
		dfs(next, cur, res)
	}
}
