package main

import (
	"unicode"
)

//420. 强密码检验器
//如果一个密码满足下述所有条件，则认为这个密码是强密码：
//由至少 6 个，至多 20 个字符组成。
//至少包含 一个小写 字母，一个大写 字母，和 一个数字 。
//同一字符 不能 连续出现三次 (比如 "...aaa..." 是不允许的, 但是"...aa...a..." 如果满足其他条件也可以算是强密码)。
//给你一个字符串 password ，返回将 password 修改到满足强密码条件需要的最少修改步数。如果 password 已经是强密码，则返回 0 。
//
//在一步修改操作中，你可以：
//
//插入一个字符到 password ，
//从 password 中删除一个字符，或
//用另一个字符来替换 password 中的某个字符。
//
//
//示例 1：
//
//输入：password = "a"
//输出：5
//示例 2：
//
//输入：password = "aA1"
//输出：3
//示例 3：
//
//输入：password = "1337C0d3"
//输出：0
//
//
//提示：
//
//1 <= password.length <= 50
//password 由字母、数字、点 '.' 或者感叹号 '!'

func strongPasswordChecker(password string) int {
	hasCap, hasLower, hasDigit := 0, 0, 0
	n := len(password)
	for _, v := range password {
		if unicode.IsLower(v) {
			hasLower = 1
		} else if unicode.IsDigit(v) {
			hasDigit = 1
		} else if unicode.IsUpper(v) {
			hasCap = 1
		}
	}
	class := hasLower + hasCap + hasDigit
	if n < 6 {
		return max(6-n, 3-class)
	} else if n <= 20 {
		cur, cnt, replace := '@', 0, 0
		for _, ch := range password {
			if ch == cur {
				cnt++
			} else {
				replace += cnt / 3
				cur = ch
				cnt = 1
			}
		}
		replace += cnt / 3
		return max(replace, 3-class)
	} else {
		replace := 0
		rmList := make([]int, 3)
		for i := 0; i < n; {
			j := i
			for j < n && password[i] == password[j] {
				j++
			}
			cnt := j - i
			if cnt >= 3 {
				replace += cnt / 3
				rmList[cnt%3]++
			}
			i = j
		}
		remove := n - 20
		cur := remove
		for i, v := range rmList {
			if i == 2 {
				v = replace
			}
			if cur != 0 && v != 0 {
				used := min(v*(i+1), cur)
				cur -= used
				replace -= used / (i + 1)
			}
		}
		return remove + max(replace, 3-class)
	}
}

func main() {
	println(strongPasswordChecker("ABABABABABABABABABAB1"))
}
