package main

//233. 数字 1 的个数
//给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
//
//示例:
//
//输入: 13
//输出: 6
//解释: 数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。

func countDigitOne(n int) int {
	result := 0
	for i := 1; i <= n; i *= 10 {
		div := i * 10
		result += n/div*i + min(max(n%div-i+1, 0), i)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(countDigitOne(999))
}
