package main

//680. 验证回文字符串 Ⅱ
//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//示例 1:
//
//输入: "aba"
//输出: True
//示例 2:
//
//输入: "abca"
//输出: True
//解释: 你可以删除c字符。
//注意:
//
//字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。

func validPalindrome(s string) bool {
	var check func(s string, l, r, k int) bool
	check = func(s string, l, r, k int) bool {
		for r > l {
			if s[l] != s[r] {
				if k > 0 && (check(s, l, r-1, k-1) || check(s, l+1, r, k-1)) {
					return true
				}
				return false
			}
			l++
			r--
		}
		return true
	}
	return check(s, 0, len(s)-1, 1)

}
