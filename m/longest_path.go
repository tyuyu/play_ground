package m

//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
//示例 1:
//
//输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//示例 2:
//
//输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func longestValidParentheses(s string) int {

	max := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] != ')' {
			continue
		}
		if s[i-1] == '(' {
			if i <= 2 {
				dp[i] = 2
			} else {
				dp[i] = dp[i-2] + 2
			}
		} else {
			if dp[i-1] > 0 {
				lf := i - 1 - dp[i-1]
				if lf >= 0 && s[lf] == '(' {
					if lf-1 > 0 {
						dp[i] = dp[lf-1] + 2 + dp[i-1]
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
