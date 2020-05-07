package str

func isScramble(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}
	set := make(map[byte]int, 0)
	for i := range s1 {
		set[s1[i]]++
	}
	for i := range s2 {
		set[s2[i]]--
	}

	for _, v := range set {
		if v != 0 {
			return false
		}
	}

	for i := 1; i < len(s1); i++ {
		t1, v1, w1 := s1[:i], s2[:i], s2[len(s2)-i:]
		t2, v2, w2 := s1[i:], s2[i:], s2[:len(s2)-i]
		if isScramble(t1, v1) && isScramble(t2, v2) {
			return true
		}
		if isScramble(t1, w1) && isScramble(t2, w2) {
			return true
		}
	}
	return false
}
