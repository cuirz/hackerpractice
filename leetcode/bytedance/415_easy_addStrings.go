package main

import "bytes"

//415. 字符串相加
//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
//注意：
//
//num1 和num2 的长度都小于 5100.
//num1 和num2 都只包含数字 0-9.
//num1 和num2 都不包含任何前导零。
//你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

func addStrings(num1 string, num2 string) string {
	n := len(num1)
	m := len(num2)
	if n < m {
		return addStrings(num2, num1)
	}
	revert := func(s string) string {
		b := bytes.Buffer{}
		size := len(s)
		for i := size - 1; i > -1; i-- {
			b.WriteByte(s[i])
		}
		return b.String()
	}
	num1 = revert(num1)
	num2 = revert(num2)
	var r byte
	result := &bytes.Buffer{}
	for i := 0; i < n; i++ {
		sum := num1[i] + r
		if i < m{
			sum += num2[i] - '0'
		}
		if sum > '0'+9 {
			result.WriteByte(sum - 10)
			r = 1
		} else {
			result.WriteByte(sum)
			r = 0
		}
	}
	if r > 0 {
		result.WriteByte('1')
	}
	return revert(result.String())

}

func main() {
	println(addStrings("90", "870"))
}
