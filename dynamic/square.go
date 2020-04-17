package dynamic

import "math"

//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
//示例 1:
//
//输入: n = 12
//输出: 3
//解释: 12 = 4 + 4 + 4.
//示例 2:
//
//输入: n = 13
//输出: 2
//解释: 13 = 4 + 9.

func numSquares(n int) int {

	if n == 0 {
		return 1
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			cut := j * j
			if cut == i {
				dp[i] = 1
				break
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-cut]+1)))
		}
	}

	return dp[n]
}
