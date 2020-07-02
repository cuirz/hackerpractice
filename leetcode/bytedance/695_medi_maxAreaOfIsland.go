package main

//695.岛屿的最大面积
//给定一个包含了一些 0 和 1 的非空二维数组 grid 。
//
//一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
//
//找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
//
//
//
//示例 1:
//
//[[0,0,1,0,0,0,0,1,0,0,0,0,0],
//[0,0,0,0,0,0,0,1,1,1,0,0,0],
//[0,1,1,0,1,0,0,0,0,0,0,0,0],
//[0,1,0,0,1,1,0,0,1,0,1,0,0],
//[0,1,0,0,1,1,0,0,1,1,1,0,0],
//[0,0,0,0,0,0,0,0,0,0,1,0,0],
//[0,0,0,0,0,0,0,1,1,1,0,0,0],
//[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。
//
//示例 2:
//
//[[0,0,0,0,0,0,0,0]]
//对于上面这个给定的矩阵, 返回 0。

func maxAreaOfIsland(grid [][]int) int {
	n := len(grid)
	if n < 1 {
		return 0
	}
	m := len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	result := 0
	dir := []int{-1, 0, 1, 0, -1}
	var search func(x, y int) int
	search = func(x, y int) int {
		res := 1
		visited[x][y] = true
		for d := 0; d < 4; d++ {
			i := x + dir[d]
			j := y + dir[d+1]
			if i >= 0 && j >= 0 && i < n && j < m && !visited[i][j] && grid[i][j] == 1 {
				res += search(i, j)
			}
		}

		return res
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				result = max(result, search(i, j))
			}

		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
