package main

import (
	"strconv"
	"unicode"
)

//640. 求解方程
//求解一个给定的方程，将x以字符串 "x=#value"的形式返回。该方程仅包含 '+' ， '-' 操作，变量x和其对应系数。
//
//如果方程没有解，请返回"No solution"。如果方程有无限解，则返回 “Infinite solutions” 。
//
//如果方程中只有一个解，要保证返回值 'x'是一个整数。
//
//
//
//示例 1：
//
//输入: equation = "x+5-3+x=6+x-2"
//输出: "x=2"
//示例 2:
//
//输入: equation = "x=x"
//输出: "Infinite solutions"
//示例 3:
//
//输入: equation = "2x=x"
//输出: "x=0"
//
//
//
//
//提示:
//
//3 <= equation.length <= 1000
//equation只有一个'='.
//equation方程由整数组成，其绝对值在[0, 100]范围内，不含前导零和变量 'x' 。

func solveEquation(equation string) string {
	sign := 1
	var x, b, i int
	n := len(equation)
	for i < n {
		if equation[i] == '=' {
			sign = -1
			i++
			continue
		}
		s := sign
		if equation[i] == '-' {
			s = -s
			i++
		} else if equation[i] == '+' {
			i++
		}
		num := 0
		hasNum := false
		for i < n && unicode.IsDigit(rune(equation[i])) {
			hasNum = true
			num = 10*num + int(equation[i]-'0')
			i++
		}
		if i < n && equation[i] == 'x' {
			if hasNum {
				s *= num
			}
			x += s
			i++
		} else {
			b += s * num
		}
	}
	if x == 0 {
		if b == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}
	if b%x != 0 {
		return "No solution"
	}
	return "x=" + strconv.Itoa(-b/x)

}
