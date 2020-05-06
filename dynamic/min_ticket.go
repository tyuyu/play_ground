package dynamic

func mincostTickets(days []int, costs []int) int {
	set := make(map[int]struct{}, 0)
	costMap := map[int]int{2: costs[0], 7: costs[1], 30: costs[2]}
	for _, v := range days {
		set[v] = struct{}{}
	}
	dp := make([]int, 366)
	for i := 1; i <= 365; i++ {
		if _, ok := set[i]; !ok {
			dp[i] = dp[i-1]
			continue
		}
		min := dp[i-1] + costMap[2]
		if i >= 7 {
			c7 := dp[i-7] + costMap[7]
			if c7 < min {
				min = c7
			}
		}
		if i >= 30 {
			c30 := dp[i-30] + costMap[30]
			if min > c30 {
				min = c30
			}
		}
		dp[i] = min
	}
	return dp[365]
}
