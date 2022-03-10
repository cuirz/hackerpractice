package main

//778. 水位上升的泳池中游泳
//在一个 N x N 的坐标方格grid 中，每一个方格的值 grid[i][j] 表示在位置 (i,j) 的平台高度。
//
//现在开始下雨了。当时间为t时，此时雨水导致水池中任意位置的水位为t。你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两个平台。假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。
//
//你从坐标方格的左上平台 (0，0) 出发。最少耗时多久你才能到达坐标方格的右下平台(N-1, N-1)？
//
//
//
//示例 1:
//
//输入: [[0,2],[1,3]]
//输出: 3
//解释:
//时间为0时，你位于坐标方格的位置为 (0, 0)。
//此时你不能游向任意方向，因为四个相邻方向平台的高度都大于当前时间为 0 时的水位。
//
//等时间到达 3 时，你才可以游向平台 (1, 1). 因为此时的水位是 3，坐标方格中的平台没有比水位 3 更高的，所以你可以游向坐标方格中的任意位置
//示例2:
//
//输入: [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
//输出: 16
//解释:
//0  1  2  3  4
//24 23 22 21  5
//12 13 14 15 16
//11 17 18 19 20
//10  9  8  7  6
//
//最终的路线用加粗进行了标记。
//我们必须等到时间为 16，此时才能保证平台 (0, 0) 和 (4, 4) 是连通的
//
//
//提示:
//
//2 <= N <= 50.
//grid[i][j] 是 [0, ..., N*N - 1] 的排列。

//思路 并查集, 最短路径
func swimInWater(grid [][]int) int {
	n := len(grid)
	parent := make([]int, n*n)
	size := make([]int, n*n)
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
	}

	type point struct {
		x, y int
	}
	pos := make([]point, n*n)
	for i, row := range grid {
		for j, h := range row {
			pos[h] = point{i, j}
		}
	}

	dir := []int{-1, 0, 1, 0, -1}
	for h := 0; ; h++ {
		for i := 0; i < 4; i++ {
			x := pos[h].x + dir[i]
			y := pos[h].y + dir[i+1]
			if x >= 0 && y >= 0 && x < n && y < n && grid[x][y] <= h {
				union(pos[h].x*n+pos[h].y, x*n+y)
			}
			if find(0) == find(n*n-1) {
				return h
			}
		}
	}
	return -1
}
