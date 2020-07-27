package main

//85. 最大矩形
//给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
//示例:
//
//输入:
//[
//["1","0","1","0","0"],
//["1","0","1","1","1"],
//["1","1","1","1","1"],
//["1","0","0","1","0"]
//]
//输出: 6

//思路 动态规划
func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n < 1 {
		return 0
	}
	m := len(matrix[0])
	dp := make([][]int, n+1)
	var result int
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i+1][j+1] = dp[i+1][j] + 1
			w := dp[i+1][j+1]
			for k := i; k > -1; k-- {
				w = min(w, dp[k+1][j+1])
				result = max(result, w*(i-k+1))
			}
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
