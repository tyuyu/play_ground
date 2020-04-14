package dynamic

import "math"

//找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。
//
//示例 1:
//
//输入:
//s = "aaabb", k = 3
//
//输出:
//3
//
//最长子串为 "aaa" ，其中 'a' 重复了 3 次。
//示例 2:
//
//输入:
//s = "ababbc", k = 2
//
//输出:
//5
//
//最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	m := make(map[rune]int, 0)
	for _, i2 := range s {
		m[i2]++
	}
	rs := []rune(s)
	bg, end := 0, len(s)
	for ; bg < len(rs); bg++ {
		if m[rs[bg]] >= k {
			break
		}
	}
	for ; end > 1; end-- {
		if m[rs[end-1]] >= k {
			break
		}
	}
	if end-bg < k {
		return 0
	}
	for i := bg; i < end; i++ {
		if m[rs[i]] < k {
			return int(math.Max(float64(longestSubstring(s[bg:i], k)), float64(longestSubstring(s[i+1:end], k))))
		}
	}
	return end - bg
}
