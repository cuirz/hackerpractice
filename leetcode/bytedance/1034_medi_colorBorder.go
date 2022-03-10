package main

//1034. 边界着色
//给你一个大小为 m x n 的整数矩阵 grid ，表示一个网格。另给你三个整数row、col 和 color 。网格中的每个值表示该位置处的网格块的颜色。
//
//当两个网格块的颜色相同，而且在四个方向中任意一个方向上相邻时，它们属于同一 连通分量 。
//
//连通分量的边界 是指连通分量中的所有与不在分量中的网格块相邻（四个方向上）的所有网格块，或者在网格的边界上（第一行/列或最后一行/列）的所有网格块。
//
//请你使用指定颜色color 为所有包含网格块grid[row][col] 的 连通分量的边界 进行着色，并返回最终的网格grid 。
//
//
//
//示例 1：
//
//输入：grid = [[1,1],[1,2]], row = 0, col = 0, color = 3
//输出：[[3,3],[3,2]]
//示例 2：
//
//输入：grid = [[1,2,2],[2,3,2]], row = 0, col = 1, color = 3
//输出：[[1,3,3],[2,3,3]]
//示例 3：
//
//输入：grid = [[1,1,1],[1,1,1],[1,1,1]], row = 1, col = 1, color = 2
//输出：[[2,2,2],[2,1,2],[2,2,2]]
//
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 50
//1 <= grid[i][j], color <= 1000
//0 <= row < m
//0 <= col < n

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[row][col] = true
	origin := grid[row][col]
	dir := []int{-1, 0, 1, 0, -1}
	array := make([][]int, 0)
	var dfs func(x, y int)
	dfs = func(x, y int) {
		isBoard := false
		for i := 0; i < 4; i++ {
			nx, ny := x+dir[i], y+dir[i+1]
			if !(nx >= 0 && ny >= 0 && nx < m && ny < n && grid[nx][ny] == origin) {
				isBoard = true
			} else if !visited[nx][ny] {
				visited[nx][ny] = true
				dfs(nx, ny)
			}
		}
		if isBoard {
			array = append(array, []int{x, y})
		}
	}
	dfs(row, col)
	for _, v := range array {
		grid[v[0]][v[1]] = color
	}

	return grid

}

func main() {
	println(colorBorder([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1, 1, 2))
}
