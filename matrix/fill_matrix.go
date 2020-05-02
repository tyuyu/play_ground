package matrix

func generateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	turn := 0
	r, l := 0, 0
	for i := 1; i <= n*n; i++ {
		switch turn % 4 {
		case 0: //go right
			matrix[r][l] = i
			if l+1 == n || matrix[r][l+1] != 0 {
				r++
				turn++
			} else {
				l++
			}
		case 1: //go down
			matrix[r][l] = i
			if r+1 == n || matrix[r+1][l] != 0 {
				l--
				turn++
			} else {
				r++
			}
		case 2: //go left
			matrix[r][l] = i
			if l-1 < 0 || matrix[r][l-1] != 0 {
				r--
				turn++
			} else {
				l--
			}
		case 3: //go up
			matrix[r][l] = i
			if matrix[r-1][l] != 0 {
				l++
				turn++
			} else {
				r--
			}
		}
	}

	return matrix
}
