package main

import "math"

//633. 平方数之和
//给定一个非负整数c，你要判断是否存在两个整数 a 和 b，使得a2 + b2 = c 。
//
//
//
//示例 1：
//
//输入：c = 5
//输出：true
//解释：1 * 1 + 2 * 2 = 5
//示例 2：
//
//输入：c = 3
//输出：false
//示例 3：
//
//输入：c = 4
//输出：true
//示例 4：
//
//输入：c = 2
//输出：true
//示例 5：
//
//输入：c = 1
//输出：true
//
//
//提示：
//
//0 <= c <= 2^31 - 1
//思路  双指针

func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		s := l*l + r*r
		if s == c {
			return true
		} else if s > c {
			r--
		} else {
			l++
		}
	}
	return false
}

func main() {
	println(judgeSquareSum(5))
}
