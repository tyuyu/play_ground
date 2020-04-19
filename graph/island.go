package graph

//岛屿数量
//给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
//
//示例 1:
//
//输入:
//11110
//11010
//11000
//00000
//
//输出: 1
//示例 2:
//
//输入:
//11000
//11000
//00100
//00011
//
//输出: 3

//不扯什么dfs什么鬼，核心思路是吧和1相邻的1都改成0，最后算1的个数
func numIslands(grid [][]byte) int {

	count := 0
	search := make([][]int, len(grid))
	for i := range search {
		search[i] = make([]int, len(grid[i]))
	}
	for i, line := range grid {
		for j := range line {
			if grid[i][j] == '1' {
				count++
				markNode(grid, i, j, search)
			}
		}
	}

	return count
}

func markNode(grid [][]byte, i int, j int, search [][]int) {

	if search[i][j] == 1 {
		return
	}
	search[i][j] = 1
	hr, hd, hl, hu := false, false, false, false
	//go right
	if j < len(grid[i])-1 {
		if grid[i][j+1] == '1' {
			markNode(grid, i, j+1, search)
			hr = true
		}
	}
	//go left
	if j > 0 {
		if grid[i][j-1] == '1' {
			markNode(grid, i, j-1, search)
			hl = true
		}
	}

	//go down
	if i < len(grid)-1 {
		if grid[i+1][j] == '1' {
			markNode(grid, i+1, j, search)
			hd = true
		}
	}
	//go up
	if i > 0 {
		if grid[i-1][j] == '1' {
			markNode(grid, i-1, j, search)
			hu = true
		}
	}

	if hr {
		grid[i][j+1] = '0'
	}
	if hd {
		grid[i+1][j] = '0'
	}
	if hl {
		grid[i][j-1] = '0'
	}
	if hu {
		grid[i-1][j] = '0'
	}
}
