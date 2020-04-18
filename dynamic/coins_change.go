package dynamic

import "sort"

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//
//
//示例 1:
//
//输入: coins = [1, 2, 5], amount = 11
//输出: 3
//解释: 11 = 5 + 5 + 1
//示例 2:
//
//输入: coins = [2], amount = 3
//输出: -1
//
//
//说明:
//你可以认为每种硬币的数量是无限的。

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return -1
	}

	dp := make([]int, amount+1)
	dp[0] = -1

	for i := 1; i <= amount; i++ {
		res := make([]int, 0)
		for _, coin := range coins {
			left := i - coin
			if left == 0 {
				res = append(res, 1)
				break
			}
			if left > 0 && dp[left] > 0 {
				res = append(res, dp[left]+1)
			}
		}
		if len(res) == 0 {
			dp[i] = -1
		} else {
			sort.Ints(res)
			dp[i] = res[0]
		}
	}
	return dp[amount]
}
