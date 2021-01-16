package main

//803. 打砖块
//有一个 m x n 的二元网格，其中 1 表示砖块，0 表示空白。砖块 稳定（不会掉落）的前提是：
//
//一块砖直接连接到网格的顶部，或者
//至少有一块相邻（4 个方向之一）砖块 稳定 不会掉落时
//给你一个数组 hits ，这是需要依次消除砖块的位置。每当消除 hits[i] = (rowi, coli) 位置上的砖块时，对应位置的砖块（若存在）会消失，然后其他的砖块可能因为这一消除操作而掉落。一旦砖块掉落，它会立即从网格中消失（即，它不会落在其他稳定的砖块上）。
//
//返回一个数组 result ，其中 result[i] 表示第 i 次消除操作对应掉落的砖块数目。
//
//注意，消除可能指向是没有砖块的空白位置，如果发生这种情况，则没有砖块掉落。
//
// 
//
//示例 1：
//
//输入：grid = [[1,0,0,0],[1,1,1,0]], hits = [[1,0]]
//输出：[2]
//解释：
//网格开始为：
//[[1,0,0,0]，
//[1,1,1,0]]
//消除 (1,0) 处加粗的砖块，得到网格：
//[[1,0,0,0]
//[0,1,1,0]]
//两个加粗的砖不再稳定，因为它们不再与顶部相连，也不再与另一个稳定的砖相邻，因此它们将掉落。得到网格：
//[[1,0,0,0],
//[0,0,0,0]]
//因此，结果为 [2] 。
//示例 2：
//
//输入：grid = [[1,0,0,0],[1,1,0,0]], hits = [[1,1],[1,0]]
//输出：[0,0]
//解释：
//网格开始为：
//[[1,0,0,0],
//[1,1,0,0]]
//消除 (1,1) 处加粗的砖块，得到网格：
//[[1,0,0,0],
//[1,0,0,0]]
//剩下的砖都很稳定，所以不会掉落。网格保持不变：
//[[1,0,0,0],
//[1,0,0,0]]
//接下来消除 (1,0) 处加粗的砖块，得到网格：
//[[1,0,0,0],
//[0,0,0,0]]
//剩下的砖块仍然是稳定的，所以不会有砖块掉落。
//因此，结果为 [0,0] 。
// 
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 200
//grid[i][j] 为 0 或 1
//1 <= hits.length <= 4 * 104
//hits[i].length == 2
//0 <= xi <= m - 1
//0 <= yi <= n - 1
//所有 (xi, yi) 互不相同

//思路  逆序思路+并查集

func hitBricks(grid [][]int, hits [][]int) []int {
	dirs := []int{-1, 0, 1, 0, -1}
	n := len(grid)
	m := len(grid[0])
	copyGrid := make([][]int, n)
	for i, v := range grid {
		copyGrid[i] = append(copyGrid[i], v...)
	}


	parents := make([]int, n*m+1)
	count := make([]int, n*m+1)

	for i := range parents {
		parents[i] = i
		count[i] = 1
	}

	for _, v := range hits {
		copyGrid[v[0]][v[1]] = 0
	}
	var find func(x int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	union := func(from, to int) {
		from, to = find(from), find(to)
		if from != to {
			count[to] += count[from]
			parents[from] = to
		}
	}

	root := n * m
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if copyGrid[i][j] == 0 {
				continue
			}
			if i == 0 {
				union(i*m+j, root)
			}
			if i > 0 && copyGrid[i-1][j] == 1 {
				union(i*m+j, (i-1)*m+j)
			}
			if j > 0 && copyGrid[i][j-1] == 1 {
				union(i*m+j, i*m+j-1)
			}
		}
	}

	result := make([]int, len(hits))
	for i := len(hits) - 1; i > -1; i-- {
		x, y := hits[i][0], hits[i][1]
		if grid[x][y] == 0 {
			continue
		}
		pre := count[find(root)]
		if x == 0 {
			union(y, root)
		}
		for k := 0; k < 4; k++ {
			a := x + dirs[k]
			b := y + dirs[k+1]
			if a >= 0 && a < n && b >= 0 && b < m && copyGrid[a][b] == 1 {
				union(x*m+y, a*m+b)
			}
		}
		post := count[find(root)]
		result[i] = max(result[i], post-pre-1)
		copyGrid[x][y] = grid[x][y]
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	hitBricks([][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}, [][]int{{1, 0}})
	//hitBricks([][]int{{1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0}, {0, 2}, {1, 1}})
}
