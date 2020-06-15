package main

import (
	"math"
	"strconv"
)

//44. 数字序列中某一位的数字
//数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
//
//请写一个函数，求任意第n位对应的数字。
//
//
//
//示例 1：
//
//输入：n = 3
//输出：3
//示例 2：
//
//输入：n = 11
//输出：0

func findNthDigitOld(n int) int {
	if n < 10 {
		return n
	}

	fn := func(i int) int {
		if i == 0 {
			return 10
		}
		return (i + 1) * 9 * int(math.Pow10(i))
	}
	x, pre := n, n
	index := 0
	for i := 0; i < 32; i++ {
		pre = x
		x -= fn(i)
		if x < 0 {
			x = pre
			index = i + 1
			break
		}
	}
	k := x % index

	result, _ := strconv.Atoi(strconv.Itoa(int(math.Pow10(index-1)) + x/index)[k : k+1])

	return result
}

func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = 9 * start * digit
	}
	num := start + (n-1)/digit
	index := (n - 1) % digit
	result, _ := strconv.Atoi(strconv.Itoa(num)[index : index+1])
	return result
}

func main() {
	println(findNthDigit(10000))
}
