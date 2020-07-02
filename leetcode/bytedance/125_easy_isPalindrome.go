package main

import "strings"

//125. 验证回文串
//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	l := 0
	r := n - 1
	for l < r {
		for l < r && !isValid(s[l]) {
			l++
		}
		for l < r && !isValid(s[r]) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isValid(a byte) bool {
	if a >= '0' && a <= '9' {
		return true
	}
	if a >= 'a' && a <= 'z' {
		return true
	}
	if a >= 'A' && a <= 'Z' {
		return true
	}
	return false

}

func main() {

}
