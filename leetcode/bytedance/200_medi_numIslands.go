package main

//200. 岛屿数量
//给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//
//
//
//示例 1:
//
//输入:
//[
//['1','1','1','1','0'],
//['1','1','0','1','0'],
//['1','1','0','0','0'],
//['0','0','0','0','0']
//]
//输出:1
//示例2:
//
//输入:
//[
//['1','1','0','0','0'],
//['1','1','0','0','0'],
//['0','0','1','0','0'],
//['0','0','0','1','1']
//]
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。

//思路

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n < 1 {
		return 0
	}
	m := len(grid[0])
	count := 0
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	dir := []int{-1, 0, 1, 0, -1}

	var search func(i, j, cnt int)
	search = func(i, j, cnt int) {
		if i < 0 || j < 0 || i >= n || j >= m || visited[i][j] || grid[i][j] == '0' {
			return
		}
		if cnt == 0 {
			count++
		}
		visited[i][j] = true
		for k := 0; k < len(dir)-1; k++ {
			search(i+dir[k], j+dir[k+1], cnt+1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			search(i, j, 0)
		}
	}
	return count

}
