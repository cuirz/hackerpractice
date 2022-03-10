package main

//343. 整数拆分
//给定一个正整数n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
//
//示例 1:
//
//输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。
//示例2:
//
//输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36。
//说明: 你可以假设n不小于 2 且不大于 58。

//思路  记忆递归
// 当 n 拆出 x,y 并记忆

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dic := make(map[int]int)
	var merge func(m int) int
	merge = func(m int) int {
		if dic[m] > 0 {
			return dic[m]
		}
		result := 0
		x, y := 0, 0
		for i := 1; i <= m>>1; i++ {
			x = merge(i)
			y = merge(m - i)
			dic[i], dic[m-i] = x, y
			result = max(result, x*y)
		}
		if result <= m {
			return m
		}
		return result
	}
	return merge(n)
}

//思路 动态规划

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i < n; i++ {
		result := 0
		for j := 1; j < i; j++ {
			result = max(result, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = result
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(integerBreak(8))
}
