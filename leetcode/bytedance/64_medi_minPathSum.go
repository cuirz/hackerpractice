package main

import "math"

//64. 最小路径和
//给定一个包含非负整数的 mxn网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
//
//示例:
//
//输入:
//[
// [1,3,1],
//[1,5,1],
//[4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。

//思路 动态规划
func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = math.MaxInt32
	}
	for j := 2; j <= m; j++ {
		dp[0][j] = math.MaxInt32
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j]) + grid[i][j]
		}
	}
	return dp[n][m]
}
