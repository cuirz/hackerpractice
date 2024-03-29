package main

//224. 基本计算器
//实现一个基本的计算器来计算一个简单的字符串表达式 s 的值。
//
//
//
//示例 1：
//
//输入：s = "1 + 1"
//输出：2
//示例 2：
//
//输入：s = " 2-1 + 2 "
//输出：3
//示例 3：
//
//输入：s = "(1+(4+5+2)-3)+(6+8)"
//输出：23
//
//
//提示：
//
//1 <= s.length <= 3* 105
//s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
//s 表示一个有效的表达式

//思路 栈
func calculate(s string) int {
	stack := []int{1}
	sign := 1
	result := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += sign * num
		}
	}
	return result
}
