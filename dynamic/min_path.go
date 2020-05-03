package dynamic

import "math"

//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
//
//示例:
//
//输入:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-path-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func minPathSum(grid [][]int) int {

	r := len(grid)
	if r == 0 {
		return 0
	}

	l := len(grid[0])
	if l == 0 {
		return 0
	}

	dp := make([][]int, r)
	for i := range dp {
		dp[i] = make([]int, l)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < l; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			left, up := math.MaxInt64, math.MaxInt64
			if j > 0 {
				left = dp[i][j-1]
			}
			if i > 0 {
				up = dp[i-1][j]
			}
			dp[i][j] = grid[i][j] + int(math.Min(float64(left), float64(up)))
		}
	}
	return dp[r-1][l-1]
}
