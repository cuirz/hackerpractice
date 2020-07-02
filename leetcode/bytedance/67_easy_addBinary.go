package main

import "strings"

//67. 二进制求和
//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
//输入为 非空 字符串且只包含数字 1 和 0。
//
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"
//
//提示：
//
//每个字符串仅由字符 '0' 或 '1' 组成。
//1 <= a.length, b.length <= 10^4
//字符串如果不是 "0" ，就都不含前导零。
//

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}

	n := len(b)
	m := len(a)
	j := m - 1
	result := make([]string, m+1)
	shift := "0"
	for i := n - 1; i > -1; i-- {
		if a[j] == b[i] {
			result[j+1] = shift
			shift = string(b[i])
		} else {
			if shift == "1" {
				result[j+1] = "0"
			} else {
				result[j+1] = "1"
			}
		}
		j--
	}

	for i := j; i > -1; i-- {
		if shift == "1" {
			if a[i] == '1' {
				result[i+1] = "0"
			} else {
				result[i+1] = shift
				shift = "0"
			}
		} else {
			result[i+1] = string(a[i])
		}
	}
	if shift == "1" {
		result[0] = "1"
		return strings.Join(result, "")
	}
	return strings.Join(result[1:], "")
}

func main() {
	println(addBinary("11", "1"))
}
