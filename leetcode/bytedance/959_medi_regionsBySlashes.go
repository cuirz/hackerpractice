package main

//959. 由斜杠划分区域
//在由 1 x 1 方格组成的 N x N 网格grid 中，每个 1 x 1方块由 /、\ 或空格构成。这些字符会将方块划分为一些共边的区域。
//
//（请注意，反斜杠字符是转义的，因此 \ 用 "\\"表示。）。
//
//返回区域的数目。
//
//提示：
//
//1 <= grid.length == grid[0].length <= 30
//grid[i][j] 是 '/'、'\'、或 ' '。

//思路  并查集
func regionsBySlashes(grid []string) int {
	n := len(grid)
	parent := make([]int, 4*n*n)
	size := make([]int, 4*n*n)
	count := 4 * n * n

	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]

	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		if size[x] < size[y] {
			x, y = y, x
		}
		size[x] += size[y]
		parent[y] = x
		count--
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			index := (i*n + j) * 4
			if i < n-1 {
				bottom := index + 4*n
				union(index+2, bottom)
			}
			if j < n-1 {
				right := index + 4 + 3
				union(index+1, right)
			}
			if grid[i][j] == '/' {
				union(index, index+3)
				union(index+1, index+2)

			} else if grid[i][j] == '\\' {
				union(index, index+1)
				union(index+2, index+3)
			} else {
				union(index, index+1)
				union(index+1, index+2)
				union(index+2, index+3)
			}
		}
	}
	return count

}
