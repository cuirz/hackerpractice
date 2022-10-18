package main

import "strconv"

//902. 最大为 N 的数字组合
//给定一个按非递减顺序排列的数字数组digits。你可以用任意次数digits[i]来写的数字。例如，如果digits = ['1','3','5']，我们可以写数字，如'13','551', 和'1351315'。
//
//返回 可以生成的小于或等于给定整数 n 的正整数的个数。
//
//
//示例 1：
//
//输入：digits = ["1","3","5","7"], n = 100
//输出：20
//解释：
//可写出的 20 个数字是：
//1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
//示例 2：
//
//输入：digits = ["1","4","9"], n = 1000000000
//输出：29523
//解释：
//我们可以写 3 个一位数字，9 个两位数字，27 个三位数字，
//81 个四位数字，243 个五位数字，729 个六位数字，
//2187 个七位数字，6561 个八位数字和 19683 个九位数字。
//总共，可以使用D中的数字写出 29523 个整数。
//示例 3:
//
//输入：digits = ["7"], n = 8
//输出：1
//
//
//提示：
//
//1 <= digits.length <= 9
//digits[i].length == 1
//digits[i]是从'1'到'9' 的数
//digits中的所有值都 不同
//digits按非递减顺序排列
//1 <= n <= 10^9

func atMostNGivenDigitSet(digits []string, n int) int {
	m := len(digits)
	s := strconv.Itoa(n)
	k := len(s)
	dp := make([][2]int, k+1)
	dp[0][1] = 1
	for i := 1; i <= k; i++ {
		for _, d := range digits {
			if d[0] == s[i-1] {
				dp[i][1] = dp[i-1][1]
			} else if d[0] < s[i-1] {
				dp[i][0] += dp[i-1][1]
			} else {
				break
			}
		}
		if i > 1 {
			dp[i][0] += m + dp[i-1][0]*m
		}
	}
	return dp[k][0] + dp[k][1]
}

func main() {
	println(atMostNGivenDigitSet([]string{"1", "2", "5", "7"}, 100))
}
