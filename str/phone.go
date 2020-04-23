package str

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//示例:
//
//输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func letterCombinations(digits string) []string {

	m := map[byte][]byte{'2': {'a', 'b', 'c'}, '3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'}, '5': {'j', 'k', 'l'}, '6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'}, '8': {'t', 'u', 'v'}, '9': {'w', 'x', 'y', 'z'},
	}

	res := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		cur := m[digits[i]]
		if len(res) == 0 {
			for _, b := range cur {
				res = append(res, string(b))
			}
		} else {
			tmp := make([]string, 0)
			for _, old := range res {
				for _, b := range cur {
					tmp = append(tmp, old+string(b))
				}
			}
			res = tmp
		}
	}
	return res
}
