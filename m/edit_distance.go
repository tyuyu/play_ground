package m

//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
//
//你可以对一个单词进行如下三种操作：
//
//插入一个字符
//删除一个字符
//替换一个字符
//
//
//示例 1：
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//示例 2：
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/edit-distance
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//官方dp
//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//5.7 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户
func minDistance(word1 string, word2 string) int {

	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			add1 := dp[i-1][j]
			add2 := dp[i][j-1]
			replace := dp[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				dp[i][j] = 1 + min(add1, add2, replace-1)
			} else {
				dp[i][j] = 1 + min(add1, add2, replace)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}
