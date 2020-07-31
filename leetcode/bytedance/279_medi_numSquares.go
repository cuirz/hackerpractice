package main

import "math"

//279. 完全平方数
//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
//示例 1:
//
//输入: n = 12
//输出: 3
//解释: 12 = 4 + 4 + 4.
//示例 2:
//
//输入: n = 13
//输出: 2
//解释: 13 = 4 + 9.

//思路 动态规划
// dp[i] = min(dp[i],dp[i-square]+1)

func numSquares(n int) int {
	array := make([]int, 0)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if n >= i*i {
			array = append(array, i*i)
			continue
		}
		break
	}
	for i:=0;i<=n;i++{
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for _, v := range array {
			if i < v {
				break
			}
			dp[i] = min(dp[i], dp[i-v]+1)
		}
	}
	return dp[n]

}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(numSquares(412))
}
