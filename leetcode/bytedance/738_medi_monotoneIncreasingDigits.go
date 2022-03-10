package main

import "strconv"

//738. 单调递增的数字
//给定一个非负整数N，找出小于或等于N的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
//
//（当且仅当每个相邻位数上的数字x和y满足x <= y时，我们称这个整数是单调递增的。）
//
//示例 1:
//
//输入: N = 10
//输出: 9
//示例 2:
//
//输入: N = 1234
//输出: 1234
//示例 3:
//
//输入: N = 332
//输出: 299
//说明: N是在[0, 10^9]范围内的一个整数。

func monotoneIncreasingDigits(N int) int {
	num := []byte(strconv.Itoa(N))
	n := len(num)
	i := 1
	for i < n && num[i-1] <= num[i] {
		i++
	}
	if i < n {
		for i > 0 && num[i-1] > num[i] {
			i--
			num[i]--
		}
		for index := i + 1; index < n; index++ {
			num[index] = '9'
		}
	}
	result, _ := strconv.Atoi(string(num))
	return result
}
