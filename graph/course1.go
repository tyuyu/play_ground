package graph

//给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
//
//
//
//示例 1:
//
//输入: 2, [[1,0]]
//输出: true
//解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
//示例 2:
//
//输入: 2, [[1,0],[0,1]]
//输出: false
//解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
//
//
//提示：
//
//输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
//你可以假定输入的先决条件中没有重复的边。
//1 <= numCourses <= 10^5

func canFinish(numCourses int, prerequisites [][]int) bool {
	arr := make([][]int, numCourses)
	for _, i := range prerequisites {
		arr[i[1]] = append(arr[i[1]], i[0])
	}
	visited := make([]int, numCourses)
	for index := 0; index < numCourses; index++ {
		if !dfs(arr, index, visited) {
			return false
		}
	}
	return true
}

func dfs(arr [][]int, index int, visited []int) bool {
	if visited[index] == 1 {
		return false
	}
	if visited[index] == -1 {
		return true
	}
	visited[index] = 1
	for _, v := range arr[index] {
		if !dfs(arr, v, visited) {
			return false
		}
	}
	visited[index] = -1
	return true
}
