package main

//397. 整数替换
//给定一个正整数n ，你可以做如下操作：
//
//如果n是偶数，则用n / 2替换n 。
//如果n是奇数，则可以用n + 1或n - 1替换n 。
//n变为 1 所需的最小替换次数是多少？
//
//
//
//示例 1：
//
//输入：n = 8
//输出：3
//解释：8 -> 4 -> 2 -> 1
//示例 2：
//
//输入：n = 7
//输出：4
//解释：7 -> 8 -> 4 -> 2 -> 1
//或 7 -> 6 -> 3 -> 2 -> 1
//示例 3：
//
//输入：n = 4
//输出：2
//
//
//提示：
//
//1 <= n <= 2^31 - 1

var mp map[int]int

func init() {
	mp = map[int]int{}
}

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if _, has := mp[n]; !has {
		if n%2 == 0 {
			mp[n] = 1 + integerReplacement(n/2)
		} else {
			mp[n] = 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))

		}
	}
	return mp[n]
}
