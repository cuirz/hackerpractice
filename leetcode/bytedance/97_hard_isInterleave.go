package main

//97. 交错字符串
//给定三个字符串s1, s2, s3, 验证s3是否是由s1和s2 交错组成的。
//
//示例 1：
//
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出：true
//示例2：
//
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出：false

//思路  动态规划
//dp[i][j]  i是 s1索引，j是s2索引
//dp[i][j] =  dp[i-1][j]  or dp[i][j-1]

func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	if n+m != len(s3) {
		return false
	}
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			index := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s3[index] == s1[i-1]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s3[index] == s2[j-1]
			}
		}
	}
	return dp[n][m]
}
