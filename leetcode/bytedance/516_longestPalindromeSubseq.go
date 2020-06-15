package main

//516. 最长回文子序列
//给定一个字符串s，找到其中最长的回文子序列。可以假设s的最大长度为1000。
//示例 1:
//输入:
//
//"bbbab"
//输出:
//
//4
//一个可能的最长回文子序列为 "bbbb"。
//示例 2:
//输入:
//
//"cbbd"
//输出:
//
//2
//一个可能的最长回文子序列为 "bb"。

// 动态规划
// dp[l][r]
// s[l] == s[r] 时 dp[l][r] = dp[l+1][r-1] + 2
// s[l] != s[r] 时 dp[l][r] = dp[l+1][r-1]
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for l := n - 1; l >= 0; l-- {
		dp[l] = make([]int, n)
		dp[l][l] = 1
		for r := l + 1; r < n; r++ {
			if s[l] == s[r] {
				dp[l][r] = dp[l+1][r-1] + 2
			} else {
				dp[l][r] = max(dp[l+1][r], dp[l][r-1])
			}
		}
	}
	return dp[0][n-1]
}
