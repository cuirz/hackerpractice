package main

//504. 七进制数
//给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
//
//
//
//示例 1:
//
//输入: num = 100
//输出: "202"
//示例 2:
//
//输入: num = -7
//输出: "-10"
//
//
//提示：
//
//-10^7 <= num <= 10^7

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var result []byte
	negative := num < 0
	if negative {
		num = -num
	}
	for num > 0 {
		result = append(result, '0'+byte(num%7))
		num /= 7
	}
	if negative {
		result = append(result, '-')
	}
	n := len(result)
	for i := 0; i < n>>1; i++ {
		result[i], result[n-i-1] = result[n-i-1], result[i]
	}
	return string(result)
}

func main() {
	println(convertToBase7(-10))
}
