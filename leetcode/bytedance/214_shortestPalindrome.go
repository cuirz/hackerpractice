package main

import "bytes"

func shortestPalindrome(s string) string {
	size := len(s)
	// 判断是否是回文
	var check func(str string) bool
	check = func(str string) bool {
		sn := len(str)
		if sn == 1 {
			return true
		}
		mid := sn / 2
		for i := 0; i < mid; i++ {
			if str[i] != str[sn-1-i] {
				return false
			}
		}
		return true
	}
	if check(s) {
		return s
	}

	var pre bytes.Buffer
	result := ""
	for j := size - 1; j > -1; j-- {
		pre.WriteByte(s[j])
		result = pre.String() + s
		if check(result) {
			return result
		}
	}
	return ""

}

func main() {
	println(shortestPalindrome("aacecaaa"))
}
