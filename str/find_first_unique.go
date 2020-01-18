package str

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//案例:
//
//s = "leetcode"
//返回 0.
//
//s = "loveleetcode",
//返回 2.
//
//
//注意事项：您可以假定该字符串只包含小写字母。

func firstUniqChar(s string) int {

	c := make([]int, 26)
	for _, r := range s {
		c[r-'a']++
	}
	for i, r := range s {
		if c[r-'a'] == 1 {
			return i
		}
	}
	return -1
}