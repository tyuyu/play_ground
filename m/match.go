package m

//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
//
//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
//两个字符串完全匹配才算匹配成功。
//
//说明:
//
//s 可能为空，且只包含从 a-z 的小写字母。
//p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
//示例 1:
//
//输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
//示例 2:
//
//输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/wildcard-matching
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isMatch(s string, p string) bool {

	i, j := 0, 0
	for j < len(p) {
		if i >= len(s) {
			if p[j] != '*' {
				return false
			}
			break
		}
		if s[i] == p[j] {
			i++
			j++
			if i < len(s) && j >= len(p) {
				return false
			}
		} else {
			if p[j] == '?' {
				i++
				j++
			} else if p[j] == '*' {
				j++
				if j >= len(p) {
					break
				}
				for s[i] != p[j] {
					i++
					if i >= len(s) {
						break
					}
				}
			} else {
				return false
			}
		}
	}

	return true
}
