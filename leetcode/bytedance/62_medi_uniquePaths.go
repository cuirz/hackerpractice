package main

//62. 不同路径
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//问总共有多少条不同的路径？
//示例1:
//
//输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右
//示例2:
//
//输入: m = 7, n = 3
//输出: 28

//思路  动态规划
func uniquePaths(m int, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n][m]
}

//思路  动态规划 压缩
func uniquePaths(m int, n int) int {
	dp := make([]int, m+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[m]
}
