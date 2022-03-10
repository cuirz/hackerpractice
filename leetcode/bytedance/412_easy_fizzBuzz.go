package main

import (
	"strconv"
	"strings"
)

//412. Fizz Buzz
//写一个程序，输出从 1 到 n 数字的字符串表示。
//
//1. 如果n是3的倍数，输出“Fizz”；
//
//2. 如果n是5的倍数，输出“Buzz”；
//
//3.如果n同时是3和5的倍数，输出 “FizzBuzz”。
//
//示例：
//
//n = 15,
//
//返回:
//[
//"1",
//"2",
//"Fizz",
//"4",
//"Buzz",
//"Fizz",
//"7",
//"8",
//"Fizz",
//"Buzz",
//"11",
//"Fizz",
//"13",
//"14",
//"FizzBuzz"
//]

func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		sb := &strings.Builder{}
		if i%3 == 0 {
			sb.WriteString("Fizz")
		}
		if i%5 == 0 {
			sb.WriteString("Buzz")
		}
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}

		result = append(result, sb.String())
	}
	return result
}
