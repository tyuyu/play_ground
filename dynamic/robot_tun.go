package dynamic

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//问总共有多少条不同的路径？
//
//
//
//例如，上图是一个7 x 3 的网格。有多少可能的路径？
//
//
//
//示例 1:
//
//输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右
//示例 2:
//
//输入: m = 7, n = 3
//输出: 28
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/unique-paths
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[n-1][m-1]
}

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
//
//
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//说明：m 和 n 的值均不超过 100。
//
//示例 1:
//
//输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: 2
//解释:
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/unique-paths-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	n := len(obstacleGrid)
	if n == 0 {
		return 1
	}
	m := len(obstacleGrid[0])
	if m == 0 {
		return 1
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				continue
			}
			//left
			left := 0
			if j > 0 {
				left = dp[i][j-1]
			}
			up := 0
			if i > 0 {
				up = dp[i-1][j]
			}
			dp[i][j] = left + up
		}
	}

	return dp[n-1][m-1]

}
