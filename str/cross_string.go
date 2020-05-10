package str

//回溯法超时了
//func isInterleave(s1 string, s2 string, s3 string) bool {
//
//	if len(s1)+len(s2) != len(s3) {
//		return false
//	}
//
//	if s1+s2 == s3 {
//		return true
//	}
//
//	if s2+s1 == s3 {
//		return true
//	}
//
//	for i := len(s3); i > 0; i-- {
//
//		pre := s3[:i]
//		if strings.HasPrefix(s1, pre) {
//			if isInterleave(strings.TrimPrefix(s1,pre),s2,strings.TrimPrefix(s3, pre)) {
//				return true
//			}
//		}
//		if strings.HasPrefix(s2, pre) {
//			if isInterleave(s1,strings.TrimPrefix(s2,pre),strings.TrimPrefix(s3, pre)) {
//				return true
//			}
//		}
//	}
//
//	return false
//}

func isInterleave(s1 string, s2 string, s3 string) bool {

	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if s1+s2 == s3 {
		return true
	}

	if s2+s1 == s3 {
		return true
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[0][0] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s3[i+j-1] == s2[j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s3[i+j-1] == s1[i-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s3[i+j-1] == s1[i-1]) || (dp[i][j-1] && s3[i+j-1] == s2[j-1])
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
