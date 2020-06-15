package main

//面试题65. 不用加减乘除做加法
//写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
//
//
//
//示例:
//
//输入: a = 1, b = 1
//输出: 2
//
//提示：
//
//a, b 均可能是负数或 0
//结果不会溢出 32 位整数

func add(a int, b int) int {
	var c int
	for b != 0 {
		c = (a & b) << 1
		a ^= b
		b = c
	}
	return a
}

func main() {
	println(add(890, 76549))
}
