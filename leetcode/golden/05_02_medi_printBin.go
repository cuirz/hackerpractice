package main

import "strings"

//05.02. 二进制数转字符串
//二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。如果该数字无法精确地用32位以内的二进制表示，则打印“ERROR”。
//
//示例1:
//
//输入：0.625
//输出："0.101"
//示例2:
//
//输入：0.1
//输出："ERROR"
//提示：0.1无法被二进制准确表示
//
//
//提示：
//
//32位包括输出中的 "0." 这两位。
//题目保证输入用例的小数位数最多只有 6 位
func printBin(num float64) string {
	str := strings.Builder{}
	str.WriteString("0.")
	for str.Len() <= 32 && num != 0 {
		num *= 2
		digit := byte(num)
		str.WriteByte('0' + digit)
		num -= float64(digit)
	}
	if str.Len() <= 32 {
		return str.String()
	}
	return "ERROR"
}
