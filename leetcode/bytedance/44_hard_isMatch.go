package main

//44. 通配符匹配
//给定一个字符串(s) 和一个字符模式(p) ，实现一个支持'?'和'*'的通配符匹配。
//
//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
//两个字符串完全匹配才算匹配成功。
//
//说明:
//
//s可能为空，且只包含从a-z的小写字母。
//p可能为空，且只包含从a-z的小写字母，以及字符?和*。
//示例1:
//
//输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
//示例2:
//
//输入:
//s = "aa"
//p = "*"
//输出: true
//解释:'*' 可以匹配任意字符串。
//示例3:
//
//输入:
//s = "cb"
//p = "?a"
//输出: false
//解释:'?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
//示例4:
//
//输入:
//s = "adceb"
//p = "*a*b"
//输出: true
//解释:第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
//示例5:
//
//输入:
//s = "acdcb"
//p = "a*c?b"
//输出: false

// 思路 动态规划

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}

	var check func(i, j int) bool
	check = func(i, j int) bool {
		if p[j-1] == '*' {
			return dp[i][j-1] || dp[i-1][j]
		}
		return dp[i-1][j-1] && (p[j-1] == '?' || s[i-1] == p[j-1])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = check(i, j)
		}
	}
	return dp[m][n]

}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	//"aaabababaaabaababbbaaaabbbbbbabbbbabbbabbaabbababab"
	//"*ab***ba**b*b*aaab*b"
	//println(isMatch("aaabababaaabaababbbaaaabbbbbbabbbbabbbabbaabbababab", "*ab***ba**b*b*aaab*b"))
	println(isMatch("a", "a*"))
}
