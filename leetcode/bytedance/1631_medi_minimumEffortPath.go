package main

import "sort"

//1631. 最小体力消耗路径
//你准备参加一场远足活动。给你一个二维rows x columns的地图heights，其中heights[row][col]表示格子(row, col)的高度。一开始你在最左上角的格子(0, 0)，且你希望去最右下角的格子(rows-1, columns-1)（注意下标从 0 开始编号）。你每次可以往 上，下，左，右四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
//
//一条路径耗费的 体力值是路径上相邻格子之间 高度差绝对值的 最大值决定的。
//
//请你返回从左上角走到右下角的最小体力消耗值。
//
//输入：heights = [[1,2,2],[3,8,2],[5,3,5]]
//输出：2
//解释：路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
//这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。
//提示：
//
//rows == heights.length
//columns == heights[i].length
//1 <= rows, columns <= 100
//1 <= heights[i][j] <= 106

//思路 并查集  二分查找 最短路径
// 二分查找，指定格子差最大值进行bfs或dfs寻路 (1~106)
// 并查集
func minimumEffortPath(heights [][]int) int {
	n := len(heights)
	m := len(heights[0])
	parent := make([]int, n*m)
	size := make([]int, n*m)
	edges := make([][]int, 0)
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
	union := func(x, y int) bool {
		x, y = find(x), find(y)
		if x == y {
			return false
		}
		if size[x] < size[y] {
			x, y = y, x
		}
		size[x] += size[y]
		parent[y] = x
		return true
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			index := i*n + j
			if i > 0 {
				edges = append(edges, []int{index - m, index, abs(heights[i][j] - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, []int{index - 1, index, abs(heights[i][j] - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	end := n*m - 1
	for _, v := range edges {
		union(v[0], v[1])
		if find(0) == find(end) {
			return v[2]
		}
	}
	return 0

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}})
}
