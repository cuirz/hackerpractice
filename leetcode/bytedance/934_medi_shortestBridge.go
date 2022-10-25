package main

//934. 最短的桥
//给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。
//
//岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。
//
//你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。
//
//返回必须翻转的 0 的最小数目。
//
//
//
//示例 1：
//
//输入：grid = [[0,1],[1,0]]
//输出：1
//示例 2：
//
//输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
//输出：2
//示例 3：
//
//输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
//输出：1
//
//
//提示：
//
//n == grid.length == grid[i].length
//2 <= n <= 100
//grid[i][j] 为 0 或 1
//grid 中恰有两个岛

func shortestBridge(grid [][]int) int {
	type pair struct {
		x, y int
	}
	dir := []int{-1, 0, 1, 0, -1}
	n := len(grid)
	m := len(grid[0])
	var step int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != 1 {
				continue
			}
			var island []pair
			grid[i][j] = -1
			queue := []pair{{i, j}}
			for len(queue) > 0 {
				top := queue[0]
				queue = queue[1:]
				island = append(island, top)
				for d := 0; d < 4; d++ {
					x, y := top.x+dir[d+0], top.y+dir[d+1]
					if 0 <= x && x < n && 0 <= y && y < m && grid[x][y] == 1 {
						queue = append(queue, pair{x, y})
						grid[x][y] = -1
					}
				}
			}
			queue = island
			for {
				tmp := queue
				queue = nil
				for _, p := range tmp {
					for d := 0; d < 4; d++ {
						x, y := p.x+dir[d+0], p.y+dir[d+1]
						if 0 <= x && x < n && 0 <= y && y < m {
							if grid[x][y] == 1 {
								return step
							}
							if grid[x][y] == 0 {
								grid[x][y] = -1
								queue = append(queue, pair{x, y})
							}
						}
					}
				}
				step++
			}
		}
	}
	return step

}
