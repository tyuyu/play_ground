package str

func wordBreaks(s string, wordDict []string) []string {
	dp := make([]bool, len(s)+1)

	dp[0] = true

	m := make(map[string]struct{}, 0)
	for _, w := range wordDict {
		m[w] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0 && !dp[i]; j-- {
			check := s[j:i]
			_, ok := m[check]
			dp[i] = dp[j] && ok
		}
	}

	if !dp[len(s)] {
		return []string{}
	}

	return dfs(dp, len(s), s, m)

}

func dfs(dp []bool, end int, s string, m map[string]struct{}) []string {
	var res []string

	if end == 0 {
		return []string{""}
	}

	for i := end - 1; i >= 0; i-- {
		if !dp[i] {
			continue
		}
		check := s[i:end]
		_, ok := m[check]
		if ok {
			for _, r := range dfs(dp, i, s, m) {
				if r != "" {
					res = append(res, r+" "+check)
				} else {
					res = append(res, check)
				}
			}
		}
	}

	return res
}
