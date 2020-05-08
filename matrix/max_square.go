package matrix

func maximalSquare(matrix [][]byte) int {

	r := len(matrix)
	if r == 0 {
		return 0
	}
	l := len(matrix[0])

	if l == 0 {
		return 0
	}

	dp := make([][]int, r)
	for i := range dp {
		dp[i] = make([]int, l)
	}

	area := 0
	for i := 0; i < r; i++ {
		for j := 0; j < l; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			a := 1
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				side := dp[i-1][j-1] + 1
				left, up := dp[i][j-1]+1, dp[i-1][j]+1
				if left < side {
					side = left
				}
				if up < side {
					side = up
				}
				dp[i][j] = side
				a = side * side
			}
			if area < a {
				area = a
			}
		}
	}
	return area

}
