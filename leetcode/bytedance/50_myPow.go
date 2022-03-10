package main

import "math"

//实现 pow(x, n) ，即计算 x 的 n 次幂函数。
//
//示例 1:
//
//输入: 2.00000, 10
//输出: 1024.00000
//示例2:
//
//输入: 2.10000, 3
//输出: 9.26100
//示例3:
//
//输入: 2.00000, -2
//输出: 0.25000
//解释: 2-2 = 1/22 = 1/4 = 0.25
//
//说明:
//
//-100.0 < x < 100.0
//n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

func myPow(x float64, n int) float64 {
	var pow func(x float64, n int) float64
	pow = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		half := pow(x, n/2)
		if n%2 == 0 {
			return half * half
		} else {
			return half * half * x
		}
	}
	if n < 0 {
		return 1 / pow(x, -n)
	}
	return pow(x, n)
}

func main() {
	println(myPow(3, -5))
	println(math.Pow(3, -5))
}
