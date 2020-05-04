package m

//45. 跳跃游戏 II
//给定一个非负整数数组，你最初位于数组的第一个位置。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
//示例:
//
//输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

//动态规则击败9.11%
//func jump(nums []int) int {
//
//	dp := make([]int, len(nums))
//
//	for i := 1; i < len(nums); i++ {
//		for j := 0; j < i; j++ {
//			if nums[j] >= i-j {
//				step := dp[j] + 1
//				if step < dp[i] || dp[i] == 0 {
//					dp[i] = step
//				}
//			}
//		}
//	}
//
//	return dp[len(nums)-1]
//}

//贪心算法 执行用时 :
//8 ms
//, 在所有 Go 提交中击败了
//95.86%
//的用户
func jump(nums []int) int {

	count := 0

	i := 0
	for i < len(nums) {
		count++
		step := nums[i]
		if i+step >= len(nums)-1 {
			return count
		}
		max := 0
		n := 0
		for j := 1; j <= step; j++ {
			next := i + j
			if nums[next]+j > max {
				max = nums[next] + j
				n = j
			}
		}
		i = i + n
	}

	return count
}
