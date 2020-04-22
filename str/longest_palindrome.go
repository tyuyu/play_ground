package str

//最长回文子串
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"

func longestPalindrome(s string) string {

	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	max, res := 0, ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPal(s, i, j, dp) {
				l := j - i + 1
				if l > max {
					max = l
					res = s[i : j+1]
				}
			}
		}
	}
	return res
}

func isPal(s string, i int, j int, dp [][]int) bool {
	if i == j {
		dp[i][j] = 1
		return true
	}
	if i+1 == j {
		if s[i] == s[j] {
			dp[i][j] = 1
			return true
		} else {
			dp[i][j] = -1
			return false
		}
	}
	if s[i] == s[j] {
		if dp[i+1][j-1] == -1 {
			dp[i][j] = -1
			return false
		} else if dp[i+1][j-1] == 1 {
			dp[i][j] = 1
			return true
		}
		b := isPal(s, i+1, j-1, dp)
		if b {
			dp[i][j] = 1
		} else {
			dp[i][j] = -1
		}
		return b
	} else {
		dp[i][j] = -1
		return false
	}
}
