package str

func minWindow(s string, t string) string {
	res := ""
	m := make(map[byte]int, 0)
	match := make(map[byte]int, 0)
	for i := range t {
		m[t[i]] = 0
		match[t[i]]++
	}
	i, j := 0, 0
	sum := 0
	for j < len(s) {
		c := s[j]
		j++
		if match[c] == 0 {
			continue
		}
		m[c]++
		if m[c] <= match[c] {
			sum++
		}
		for sum == len(t) {
			tmp := s[i:j]
			if res == "" || len(tmp) < len(res) {
				res = tmp
			}
			l := s[i]
			i++
			if _, ok := match[l]; !ok {
				continue
			}
			if m[l] == match[l] {
				sum--
			}
			m[l]--
		}
	}
	return res
}
