package main

//227. 基本计算器 II
//给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//
//整数除法仅保留整数部分。
//
// 
//
//示例 1：
//
//输入：s = "3+2*2"
//输出：7
//示例 2：
//
//输入：s = " 3/2 "
//输出：1
//示例 3：
//
//输入：s = " 3+5 / 2 "
//输出：5
// 
//
//提示：
//
//1 <= s.length <= 3 * 105
//s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
//s 表示一个 有效表达式
//表达式中的所有整数都是非负整数，且在范围 [0, 231 - 1] 内
//题目数据保证答案是一个 32-bit 整数

func calculate(s string) int {
	stack := make([]int, 0)
	sign := '+'
	num := 0
	for i, c := range s {
		isDigit := c >= '0' && c <= '9'
		if isDigit {
			num = num*10 + int(c-'0')
		}
		if (!isDigit && s[i] != ' ') || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			sign = c
			num = 0
		}
	}
	var result int
	for _, v := range stack {
		result += v
	}
	return result
}
