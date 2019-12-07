package str

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false

func isPalindrome(s string) bool {

	if s == "" {
		return true
	}

	i, j := 0, len(s)-1
	for i < j {
		for !isNum(s[i]) && !isLetters(s[i]) {
			i++
			if i == j {
				return true
			}
		}
		for !isNum(s[j]) && !isLetters(s[j]) {
			j--
		}
		if isNum(s[i]) && s[i] != s[j] {
			return false
		}
		diff := int(s[i]) - int(s[j])
		if diff != 0 && diff != 32 && diff != -32 {
			return false
		}
		i++
		j--
	}

	return true
}

func isNum(v byte) bool {
	return v >= 48 && v <= 57
}

func isLetters(v byte) bool {
	return (v >= 65 && v <= 90) || (v >= 97 && v <= 122)
}
