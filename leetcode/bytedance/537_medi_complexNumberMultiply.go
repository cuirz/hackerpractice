package main

import (
	"fmt"
	"strconv"
	"strings"
)

//537. 复数乘法
//复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：
//
//实部 是一个整数，取值范围是 [-100, 100]
//虚部 也是一个整数，取值范围是 [-100, 100]
//i2 == -1
//给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。
//
//
//
//示例 1：
//
//输入：num1 = "1+1i", num2 = "1+1i"
//输出："0+2i"
//解释：(1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
//示例 2：
//
//输入：num1 = "1+-1i", num2 = "1+-1i"
//输出："0+-2i"
//解释：(1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。
//
//
//提示：
//
//num1 和 num2 都是有效的复数表示。
func complexNumberMultiply(num1 string, num2 string) string {
	parse := func(s string) (int, int) {
		i := strings.IndexByte(s, '+')
		real, _ := strconv.Atoi(s[:i])
		img, _ := strconv.Atoi(s[i+1 : len(s)-1])
		return real, img
	}
	a1, b1 := parse(num1)
	a2, b2 := parse(num2)
	return fmt.Sprintf("%d+%di", a1*a2-b1*b2, a1*b2+a2*b1)
}
