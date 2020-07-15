package main

import "math"

//120. 三角形最小路径和
//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
//相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//
//
//例如，给定三角形：
//
//[
//[2],
//[3,4],
//[6,5,7],
//[4,1,8,3]
//]
//自顶向下的最小路径和为11（即，2+3+5+1= 11）。
//
//
//说明：
//
//如果你可以只使用 O(n)的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

//思路1 dfs暴力
//思路2 动态规划，倒序
//因为三角结构数组，最底层数组数=三角层数
//dp[n-2][0]=min(dp[n-1][0],dp[n-1][1]) + triangle[n-2][0]
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	for i := n - 1; i > -1; i-- {
		m := len(triangle[i])
		for j := 0; j < m; j++ {
			if i == n-1 {
				dp[j] = triangle[i][j]
			} else {
				dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0]
}
func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	var dfs func(dep, index int) int
	dfs = func(dep, index int) int {
		if dep == n {
			return 0
		}
		if index >= len(triangle[dep]) {
			return math.MaxInt32
		}
		return min(triangle[dep][index]+dfs(dep+1, index), triangle[dep][index]+dfs(dep+1, index+1))
	}
	return min(triangle[0][0]+dfs(1, 0), triangle[0][0]+dfs(1, 1))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
