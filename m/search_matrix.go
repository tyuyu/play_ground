package m

func searchMatrix(matrix [][]int, target int) bool {
	i := findLine(matrix, 0, len(matrix)-1, target)
	if i == -1 {
		return false
	}
	lines := matrix[i]
	return findTarget(lines, 0, len(lines)-1, target)
}

func findTarget(lines []int, bg int, end int, target int) bool {
	mid := (bg + end) / 2
	if lines[mid] == target {
		return true
	}
	if bg == end {
		return false
	}
	if target > lines[mid] {
		return findTarget(lines, mid, end, target)
	}
	return findTarget(lines, bg, mid, target)
}

func findLine(matrix [][]int, bg int, end int, target int) int {

	mid := (bg + end) / 2
	if matrix[mid][0] <= target && matrix[mid][len(matrix)-1] >= target {
		return mid
	}
	if end == bg {
		return -1
	}
	if target > matrix[mid][len(matrix)-1] {
		return findLine(matrix, mid, end, target)
	}
	return findLine(matrix, bg, mid, target)

}
