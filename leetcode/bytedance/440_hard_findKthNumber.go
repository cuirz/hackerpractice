package main

//440. 字典序的第K小数字
//给定整数n和k，返回[1, n]中字典序第k小的数字。
//
//
//
//示例 1:
//
//输入: n = 13, k = 2
//输出: 10
//解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
//示例 2:
//
//输入: n = 1, k = 1
//输出: 1
//
//
//提示:
//
//1 <= k <= n <= 10^9

func findKthNumber(n int, k int) int {
	getChildren := func(value int) (children int) {
		first, last := value, value
		for first <= n {
			children += min(last, n) - first + 1
			first *= 10
			last = last*10 + 9
		}
		return
	}
	result := 1
	k -= 1
	for k > 0 {
		children := getChildren(result)
		if children <= k {
			k -= children
			result++
		} else {
			result *= 10
			k--
		}
	}
	return result
}
