package main

import "bytes"

//917. 仅仅反转字母

//给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
//示例 1：
//
//输入："ab-cd"
//输出："dc-ba"
//示例 2：
//
//输入："a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba"
//示例 3：
//
//输入："Test1ng-Leet=code-Q!"
//输出："Qedo1ct-eeLg=ntse-T!"
//
//
//提示：
//
//S.length <= 100
//33 <= S[i].ASCIIcode <= 122
//S 中不包含 \ or "

func reverseOnlyLetters(S string) string {
	n := len(S)
	var result bytes.Buffer
	index := n - 1

	for i := 0; i < n; i++ {
		if !valid(S[i]) {
			result.WriteByte(S[i])
		} else {
			for !valid(S[index]) {
				index--
			}
			result.WriteByte(S[index])
			index--
		}
	}
	return result.String()
}
func valid(x byte) bool {
	if x >= 'a' && x <= 'z' {
		return true
	} else if x >= 'A' && x <= 'Z' {
		return true
	}
	return false

}

func main() {
	println(reverseOnlyLetters("=09zZa-bC-dEf-ghIj"))
}
