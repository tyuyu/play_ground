package str

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//说明:
//你可以假设字符串只包含小写字母。
//
//进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

//考虑通用性，直接用rune + map实现

//不考虑通用性，数组计数可以更快点

func isAnagram(s string, t string) bool {

	if s == t {
		return false
	}
	if len(s) != len(t) {
		return false
	}

	sm, tm := make(map[rune]int, 0), make(map[rune]int, 0)

	for _, r := range s {
		if x, ok := sm[r]; ok {
			sm[r] = x + 1
		} else {
			sm[r] = 1
		}
	}
	for _, r := range t {
		if x, ok := tm[r]; ok {
			tm[r] = x + 1
		} else {
			tm[r] = 1
		}
	}

	for r, c := range sm {
		sc, ok := tm[r]
		if !ok {
			return false
		}
		if sc != c {
			return false
		}
	}

	return true
}
