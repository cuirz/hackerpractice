package main

//131. 分割回文串
//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
//回文串 是正着读和反着读都一样的字符串。
//
//
//
//示例 1：
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//示例 2：
//
//输入：s = "a"
//输出：[["a"]]
//
//
//提示：
//
//1 <= s.length <= 16
//s 仅由小写英文字母组成

//思路  回溯+动态规划
func partition(s string) [][]string {
	result := make([][]string, 0)
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}
	splits := make([]string, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			result = append(result, append([]string{}, splits...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]

			}
		}

	}
	dfs(0)
	return result

}

func main() {
	res := partition("aab")
	print(res)
}
