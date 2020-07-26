package dynamic

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

	ans := make([][]int, len(matrix))
	for i := range ans {
		ans[i] = make([]int, len(matrix[0]))
	}

	max := 1
	for i := range matrix {
		for j := range matrix[i] {
			v := calPath(i, j, ans, matrix)
			if v > max {
				max = v
			}
		}
	}

	return max
}

func calPath(i, j int, ans, matrix [][]int) int {

	if ans[i][j] > 0 {
		return ans[i][j]
	}
	cur := matrix[i][j]

	left, right, up, down := 0, 0, 0, 0
	if i > 0 && matrix[i-1][j] > cur {
		up = calPath(i-1, j, ans, matrix)
	}

	if i+1 < len(matrix) && matrix[i+1][j] > cur {
		down = calPath(i+1, j, ans, matrix)
	}

	if j > 0 && matrix[i][j-1] > cur {
		left = calPath(i, j-1, ans, matrix)
	}

	if j+1 < len(matrix[i]) && matrix[i][j+1] > cur {
		right = calPath(i, j+1, ans, matrix)
	}

	max := 0
	for _, v := range []int{up, down, left, right} {
		if v > max {
			max = v
		}
	}
	ans[i][j] = max + 1
	return max + 1
}
