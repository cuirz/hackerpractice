package swordoffer

import "strings"

//剑指 Offer 20. 表示数值的字符串
//请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。

func isNumber(s string) bool {
	var mark, point, number, exp int
	s = strings.TrimSpace(s)
	for _, v := range s {
		switch v {
		case '+', '-':
			if mark > 0 || number > 0 {
				return false
			}
			mark++
		case 'e', 'E':
			if number == 0 || exp > 0 {
				return false
			}
			exp++
			mark = 0
			number = 0
			point = 0
		case '.':
			if exp > 0 || point > 0 {
				return false
			}
			point++
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number++
		case ' ':
			return false
		default:
			return false
		}
	}
	return !((exp > 0 || mark > 0 || point > 0) && number == 0) && number > 0
}
