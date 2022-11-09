package main

//764. 最大加号标志
//在一个 n x n 的矩阵 grid 中，除了在数组 mines 中给出的元素为 0，其他每个元素都为 1。mines[i] = [xi, yi]表示 grid[xi][yi] == 0
//
//返回  grid 中包含 1 的最大的 轴对齐 加号标志的阶数 。如果未找到加号标志，则返回 0 。
//
//一个 k 阶由 1 组成的 “轴对称”加号标志 具有中心网格 grid[r][c] == 1 ，以及4个从中心向上、向下、向左、向右延伸，长度为 k-1，由 1 组成的臂。注意，只有加号标志的所有网格要求为 1 ，别的网格可能为 0 也可能为 1 。
//
//
//
//示例 1：
//
//
//
//输入: n = 5, mines = [[4, 2]]
//输出: 2
//解释: 在上面的网格中，最大加号标志的阶只能是2。一个标志已在图中标出。
//示例 2：
//
//
//
//输入: n = 1, mines = [[0, 0]]
//输出: 0
//解释: 没有加号标志，返回 0 。
//
//
//提示：
//
//1 <= n <= 500
//1 <= mines.length <= 5000
//0 <= xi, yi < n
//每一对 (xi, yi) 都 不重复
func orderOfLargestPlusSign(n int, mines [][]int) int {
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, n)
	}
	grid := make(map[int]bool)
	for _, v := range mines {
		grid[v[0]*n+v[1]] = true
	}

	left, top := 0, 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			beforeLeft, beforeTop := 0, 0
			if j > 0 {
				beforeLeft = dp[i][j-1][left]
			}
			if i > 0 {
				beforeTop = dp[i-1][j][top]
			}
			if grid[i*n+j] {
				dp[i][j][left] = 0
				dp[i][j][top] = 0
			} else {
				dp[i][j][left] = beforeLeft + 1
				dp[i][j][top] = beforeTop + 1
			}
		}
	}

	right, bottom := 0, 1
	var result int
	for i := n - 1; i > -1; i-- {
		for j := n - 1; j > -1; j-- {
			afterRight, afterBottom := 0, 0
			if i < n-1 {
				afterBottom = dp[i+1][j][bottom]
			}
			if j < n-1 {
				afterRight = dp[i][j+1][right]
			}
			if grid[i*n+j] {
				dp[i][j][right] = 0
				dp[i][j][bottom] = 0
			} else {
				afterRight += 1
				afterBottom += 1
				m := min(dp[i][j][left], dp[i][j][top], afterRight, afterBottom)
				if m > result {
					result = m
				}
				dp[i][j][right] = afterRight
				dp[i][j][bottom] = afterBottom
			}
		}
	}
	return result
}

func min(x ...int) int {
	res := x[0]
	for _, v := range x[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
