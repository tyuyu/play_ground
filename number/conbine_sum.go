package number

import "sort"

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
//
//说明：
//
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。
//示例 1:
//
//输入: candidates = [2,3,6,7], target = 7,
//所求解集为:
//[
//  [7],
//  [2,2,3]
//]
//示例 2:
//
//输入: candidates = [2,3,5], target = 8,
//所求解集为:
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//]
//

func combinationSum(candidates []int, target int) [][]int {

	res := make([][]int, 0)
	sort.Ints(candidates)

	dfs(candidates, []int{}, target, 0, &res)

	return res
}

func dfs(candidates []int, last []int, target, index int, res *[][]int) {

	if target == 0 {
		*res = append(*res, last)
	}

	for i := index; i < len(candidates); i++ {
		num := candidates[i]
		if num > target {
			return
		}
		next := make([]int, 0)
		next = append(next, last...)
		next = append(next, num)
		dfs(candidates, next, target-num, i, res)
	}
}

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的每个数字在每个组合中只能使用一次。
//
//说明：
//
//所有数字（包括目标数）都是正整数。
//解集不能包含重复的组合。
//示例 1:
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
//示例 2:
//
//输入: candidates = [2,5,2,1,2], target = 5,
//所求解集为:
//[
//  [1,2,2],
//  [5]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combination-sum-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := new([][]int)
	dsf2(candidates, []int{}, target, res)
	return *res
}

func dsf2(candidates []int, old []int, target int, res *[][]int) {

	if target == 0 {
		*res = append(*res, old)
		return
	}

	for i, num := range candidates {
		if num > target {
			return
		}
		if i > 0 && num == candidates[i-1] {
			continue
		}
		nc := candidates[i+1:]
		next := append(old, num)
		dsf2(nc, next, target-num, res)
	}
}
