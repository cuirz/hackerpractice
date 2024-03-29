package main

//63. 不同路径 II
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//说明：m和 n 的值均不超过 100。
//
//示例1:
//
//输入:
//[
// [0,0,0],
// [0,1,0],
// [0,0,0]
//]
//输出: 2
//解释:
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右

//思路 动态规划
// dp[i][j] = dp[i-1][j] +dp[i][j-1]

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n < 1 {
		return 0
	}
	m := len(obstacleGrid[0])
	if m < 1 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[1][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if obstacleGrid[i-1][j-1] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}
