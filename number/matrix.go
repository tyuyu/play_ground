package number

//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
//给定 target = 5，返回 true。
//给定 target = 20，返回 false。

func searchMatrix(matrix [][]int, target int) bool {

	i, j := 0, 0
	for {
		if i >= len(matrix) {
			break
		}
		array := matrix[i]
		if j >= len(array) {
			break
		}
		now := array[j]
		if now == target {
			return true
		}
		//go right
		if j+1 < len(array) && array[j+1] <= target {
			j++
			continue
		}
		//go left
		if j-1 >= 0 && array[j-1] >= target {
			j--
			continue
		}
		i++
	}

	return false
}
