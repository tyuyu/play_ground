package dynamic

import "math"

//最长连续序列
//给定一个未排序的整数数组，找出最长连续序列的长度。
//
//要求算法的时间复杂度为 O(n)。
//
//示例:
//
//输入: [100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
	}

	max := 0
	for _, num := range nums {
		count := 1
		left, right := num-1, num+1
		_, ok := m[left]
		for ok {
			delete(m, left)
			left--
			count++
			_, ok = m[left]
		}
		_, ok = m[right]
		for ok {
			delete(m, right)
			right++
			count++
			_, ok = m[right]
		}
		if count > max {
			max = count
		}
	}

	return max
}

//矩阵中的最长递增路径
//给定一个整数矩阵，找出最长递增路径的长度。
//
//对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。
//
//示例 1:
//
//输入: nums =
//[
//  [9,9,4],
//  [6,6,8],
//  [2,1,1]
//]
//输出: 4
//解释: 最长递增路径为 [1, 2, 6, 9]。
//示例 2:
//
//输入: nums =
//[
//  [3,4,5],
//  [3,2,6],
//  [2,2,1]
//]
//输出: 4
//解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。

func longestIncreasingPath(matrix [][]int) int {

	if len(matrix) == 0 {
		return 0
	}

	if len(matrix[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}

	max := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if dp[i][j] != -1 {
				if max < dp[i][j] {
					max = dp[i][j]
				}
				continue
			}
			v := dsf(matrix, i, j, dp)
			if max < v {
				max = v
			}
		}
	}

	return max
}

func dsf(matrix [][]int, i int, j int, dp [][]int) int {

	max := 1
	//go up
	if i > 0 && matrix[i][j] < matrix[i-1][j] {
		if dp[i-1][j] != -1 {
			max = dp[i-1][j] + 1
		} else {
			max = dsf(matrix, i-1, j, dp) + 1
		}
	}
	//go down
	if i < len(matrix)-1 && matrix[i][j] < matrix[i+1][j] {
		if dp[i+1][j] != -1 {
			max = int(math.Max(float64(dp[i+1][j]+1), float64(max)))
		} else {
			max = int(math.Max(float64(dsf(matrix, i+1, j, dp)+1), float64(max)))
		}
	}
	//go left
	if j > 0 && matrix[i][j] < matrix[i][j-1] {
		if dp[i][j-1] != -1 {
			max = int(math.Max(float64(dp[i][j-1]+1), float64(max)))
		} else {
			max = int(math.Max(float64(dsf(matrix, i, j-1, dp)+1), float64(max)))
		}
	} //go right
	if j < len(matrix[0])-1 && matrix[i][j] < matrix[i][j+1] {
		if dp[i][j+1] != -1 {
			max = int(math.Max(float64(dp[i][j+1]+1), float64(max)))
		} else {
			max = int(math.Max(float64(dsf(matrix, i, j+1, dp)+1), float64(max)))
		}
	}

	dp[i][j] = max
	return max
}
