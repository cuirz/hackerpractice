package main

import "math"

//1293. 网格中的最短路径
//给你一个m * n的网格，其中每个单元格不是0（空）就是1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。
//如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。如果找不到这样的路径，则返回 -1。

func shortestPathBfs(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}
	x, y := 0, 0
	visited := make([][][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][]bool, m)
		for j := 0; j < m; j++ {
			visited[i][j] = make([]bool, k+1)
		}
	}

	stack := make([]pos, 0)
	stack = append(stack, pos{x: 0, y: 0, rest: k})

	dir := []int{-1, 0, 1, 0, -1}
	index := 0
	count := 0

	for len(stack) > 0 {
		count++
		size := len(stack)
		for j := 0; j < size; j++ {
			p := stack[0]
			stack = stack[1:]
			for i := 0; i < 4; i++ {
				//判断越界
				x = p.x + dir[i]
				y = p.y + dir[i+1]
				if x < 0 || y < 0 || x > m-1 || y > n-1 {
					continue
				}

				index = k - p.rest

				if x == m-1 && y == n-1 {
					return count
				}

				if grid[y][x] == 0 && !visited[y][x][index] {
					visited[y][x][index] = true
					stack = append(stack, pos{x: x, y: y, rest: p.rest})
				} else if grid[y][x] == 1 && p.rest > 0 && !visited[y][x][index+1] {
					visited[y][x][index+1] = true
					stack = append(stack, pos{x: x, y: y, rest: p.rest - 1})
				}
			}

		}

	}
	return -1
}

type pos struct {
	x    int
	y    int
	rest int
}

func shortestPath(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}
	if k >= m+n-3 {
		return m + n - 2
	}
	const maxsize = 40
	var visited [maxsize][maxsize]bool

	var dfs func(x, y, rest, step int) int
	dfs = func(x, y, rest, step int) int {
		if x < 0 || y < 0 || x >= m || y >= n || visited[y][x] || grid[y][x] == 1 && rest == 0 {
			return math.MaxInt32
		}
		if x == m-1 && y == n-1 {
			return step
		}
		if grid[y][x] == 1 {
			rest--
		}
		visited[y][x] = true
		step++
		left := dfs(x-1, y, rest, step)
		right := dfs(x+1, y, rest, step)
		up := dfs(x, y-1, rest, step)
		down := dfs(x, y+1, rest, step)
		visited[y][x] = false
		return min(min(left, right), min(up, down))
	}
	result := dfs(0, 0, k, 0)
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	//println(shortestPath([][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}}, 1))
	println(shortestPath([][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}, 1))
	//[[0,1,1],[1,1,1],[1,0,0]]
	//1
}
