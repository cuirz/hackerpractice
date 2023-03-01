package main

//2373. 矩阵中的局部最大值
//给你一个大小为 n x n 的整数矩阵 grid 。
//
//生成一个大小为 (n - 2) x (n - 2) 的整数矩阵  maxLocal ，并满足：
//
//maxLocal[i][j] 等于 grid 中以 i + 1 行和 j + 1 列为中心的 3 x 3 矩阵中的 最大值 。
//换句话说，我们希望找出 grid 中每个 3 x 3 矩阵中的最大值。
//
//返回生成的矩阵。
//
//
//
//示例 1：
//
//
//
//输入：grid = [[9,9,8,1],[5,6,2,6],[8,2,6,4],[6,2,2,2]]
//输出：[[9,9],[8,6]]
//解释：原矩阵和生成的矩阵如上图所示。
//注意，生成的矩阵中，每个值都对应 grid 中一个相接的 3 x 3 矩阵的最大值。
//示例 2：
//
//
//
//输入：grid = [[1,1,1,1,1],[1,1,1,1,1],[1,1,2,1,1],[1,1,1,1,1],[1,1,1,1,1]]
//输出：[[2,2,2],[2,2,2],[2,2,2]]
//解释：注意，2 包含在 grid 中每个 3 x 3 的矩阵中。
//
//
//提示：
//
//n == grid.length == grid[i].length
//3 <= n <= 100
//1 <= grid[i][j] <= 100
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	result := make([][]int, n-2)
	biggest := func(i, j int) int {
		best := grid[i][j]
		for r := i - 1; r <= i+1; r++ {
			for c := j - 1; c <= j+1; c++ {
				if grid[r][c] > best {
					best = grid[r][c]
				}
			}
		}
		return best
	}
	for i := 1; i < n-1; i++ {
		row := make([]int, n-2)
		for j := 1; j < n-1; j++ {
			best := biggest(i, j)
			row[j-1] = best
		}
		result[i-1] = row
	}
	return result
}
