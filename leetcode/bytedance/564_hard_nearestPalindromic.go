package main

import (
	"math"
	"strconv"
)

//564. 寻找最近的回文数
//给定一个表示整数的字符串 n ，返回与它最近的回文整数（不包括自身）。如果不止一个，返回较小的那个。
//
//“最近的”定义为两个整数差的绝对值最小。
//
//
//
//示例 1:
//
//输入: n = "123"
//输出: "121"
//示例 2:
//
//输入: n = "1"
//输出: "0"
//解释: 0 和 2是最近的回文，但我们返回最小的，也就是 0。
//
//
//提示:
//
//1 <= n.length <= 18
//n 只由数字组成
//n 不含前导 0
//n 代表在 [1, 10^18 - 1] 范围内的整数

func nearestPalindromic(n string) string {
	l := len(n)
	if l == 1 {
		num, _ := strconv.Atoi(n)
		return strconv.Itoa(num - 1)
	}
	result := []int{int(math.Pow10(l-1)) - 1, int(math.Pow10(l) + 1)}
	half, _ := strconv.Atoi(n[:(l+1)/2])
	for _, num := range []int{half - 1, half, half + 1} {
		x := num
		if l&1 == 1 {
			x /= 10
		}
		for ; x > 0; x /= 10 {
			num = num*10 + x%10
		}
		result = append(result, num)
	}
	origin, _ := strconv.Atoi(n)
	ans := -1
	for _, i := range result {
		if i != origin {
			if ans == -1 || abs(i-origin) < abs(ans-origin) || abs(i-origin) == abs(ans-origin) && i < ans {
				ans = i
			}
		}
	}
	return strconv.Itoa(ans)
}
