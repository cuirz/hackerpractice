package main

//263. 丑数
//给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
//
//丑数 就是只包含质因数2、3 和/或5的正整数。
//
//
//
//示例 1：
//
//输入：n = 6
//输出：true
//解释：6 = 2 × 3
//示例 2：
//
//输入：n = 8
//输出：true
//解释：8 = 2 × 2 × 2
//示例 3：
//
//输入：n = 14
//输出：false
//解释：14 不是丑数，因为它包含了另外一个质因数7 。
//示例 4：
//
//输入：n = 1
//输出：true
//解释：1 通常被视为丑数。

func isUgly(n int) bool {
	if n == 1 {
		return true
	}
	for _, v := range []int{2, 3, 5} {
		d := n / v
		r := n % v
		if d > 0 && r == 0 {
			return isUgly(d)
		}
	}
	return false
}
