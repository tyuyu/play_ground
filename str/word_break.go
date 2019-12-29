package str

import "log"

func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)
	so := make([]string, len(s)+1)

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
			if dp[i] {
				so[i] = so[j] + check + ","
			}
		}
	}
	if dp[len(s)] {
		log.Print("solute is " + so[len(s)])
	}
	return dp[len(s)]
}
