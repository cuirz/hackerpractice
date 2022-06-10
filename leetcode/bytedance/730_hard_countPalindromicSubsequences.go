package main

//730. 统计不同回文子序列
//给定一个字符串 s，返回 s中不同的非空「回文子序列」个数 。
//
//通过从 s中删除 0 个或多个字符来获得子序列。
//
//如果一个字符序列与它反转后的字符序列一致，那么它是「回文字符序列」。
//
//如果有某个 i , 满足ai!= bi，则两个序列a1, a2, ...和b1, b2, ...不同。
//
//注意：
//
//结果可能很大，你需要对10^9+ 7取模 。
//
//
//示例 1：
//
//输入：s = 'bccb'
//输出：6
//解释：6 个不同的非空回文子字符序列分别为：'b', 'c', 'bb', 'cc', 'bcb', 'bccb'。
//注意：'bcb' 虽然出现两次但仅计数一次。
//示例 2：
//
//输入：s = 'abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba'
//输出：104860361
//解释：共有 3104860382 个不同的非空回文子序列，104860361 对 10^9 + 7 取模后的值。
//
//
//提示：
//
//1 <= s.length <= 1000
//s[i]仅包含'a','b','c'或'd'

func countPalindromicSubsequences(s string) int {
	const mod int = 1e9 + 7
	n := len(s)
	var dp [4][][]int
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	for i, c := range s {
		dp[c-'a'][i][i] = 1
	}
	ch := []byte{'a', 'b', 'c', 'd'}
	for sz := 2; sz <= n; sz++ {
		for i, j := 0, sz-1; j < n; i++ {
			for k := 0; k < 4; k++ {
				if s[i] == ch[k] && s[j] == ch[k] {
					dp[k][i][j] = (2 + dp[0][i+1][j-1] + dp[1][i+1][j-1] + dp[2][i+1][j-1] + dp[3][i+1][j-1]) % mod
				} else if s[i] == ch[k] {
					dp[k][i][j] = dp[k][i][j-1]
				} else if s[j] == ch[k] {
					dp[k][i][j] = dp[k][i+1][j]
				} else {
					dp[k][i][j] = dp[k][i+1][j-1]
				}
			}
			j++
		}
	}
	var result int
	for i := 0; i < 4; i++ {
		result += dp[i][0][n-1]
	}
	return result % mod

}
