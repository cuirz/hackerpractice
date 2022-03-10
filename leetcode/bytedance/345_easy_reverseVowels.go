package main

import "strings"

//345. 反转字符串中的元音字母
//编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
//
//
//
//示例 1：
//
//输入："hello"
//输出："holle"
//示例 2：
//
//输入："leetcode"
//输出："leotcede"
//
//
//提示：
//
//元音字母不包含字母 "y" 。

func reverseVowels(s string) string {
	n := len(s)
	result := []byte(s)
	l, r := 0, n-1
	isVow := func(b string) bool {
		return strings.Contains("aeiouAEIOU", b)
	}
	for l < r {
		for l < n && !isVow(string(s[l])) {
			l++
		}
		for r >= 0 && !isVow(string(s[r])) {
			r--
		}
		if l < r {
			result[l], result[r] = result[r], result[l]
			l++
			r--
		}
	}
	return string(result)
}
func main() {
	println(reverseVowels("hello"))
}
