package main

import (
	"strconv"
	"strings"
)

//43. 字符串相乘
//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//示例 1:
//
//输入: num1 = "2", num2 = "3"
//输出: "6"
//示例 2:
//
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//说明：
//
//num1 和 num2 的长度小于110。
//num1 和 num2 只包含数字 0-9。
//num1 和 num2 均不以零开头，除非是数字 0 本身。
//不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

func multiply(num1 string, num2 string) string {
	n := len(num1)
	m := len(num2)
	if n < m {
		return multiply(num2, num1)
	}
	res := make([]byte, n+m+1)
	var x, y byte
	for i := n - 1; i > -1; i-- {
		for j := m - 1; j > -1; j-- {
			index := (n - i - 1) + (m - j - 1)
			x = num1[i] - '0'
			y = num2[j] - '0'
			muti := x * y
			res[index] += muti % 10
			res[index+1] += muti/10 + res[index]/10
			res[index+2] += res[index+1] / 10
			res[index] = res[index] % 10
			res[index+1] = res[index+1] % 10
		}
	}
	result := &strings.Builder{}
	skip := true
	for i := len(res) - 1; i > -1; i-- {
		if skip && res[i] > 0 {
			skip = false
		}
		if !skip {
			result.WriteString(strconv.Itoa(int(res[i])))
		}
	}
	if result.Len() == 0 {
		return "0"
	}
	return result.String()

}

func main() {
	s := multiply("0", "0")

	println(s)
}
