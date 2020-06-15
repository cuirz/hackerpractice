package main

import "strconv"

//面试题46. 把数字翻译成字符串
//给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
//
//
//示例 1:
//
//输入: 12258
//输出: 5
//解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

//思路： 动态规划
// dp[s] = dp[s-1] + dp[s-2]
func translateNum(num int) int {
	str := strconv.Itoa(num)
	var dp func(s string) int
	dp = func(s string) int {
		n := len(s)
		if n < 2 {
			return 1
		}
		if s[0] == '1' || (s[0] == '2' && s[1] < '6') {
			return dp(s[1:]) + dp(s[2:])
		}
		return dp(s[1:])
	}
	return dp(str)
}

func main() {
	println(translateNum(18822))
	println(translateNum(220))
	println(translateNum(22220))
}
