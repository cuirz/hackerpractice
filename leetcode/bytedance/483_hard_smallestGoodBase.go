package main

import (
	"math"
	"math/bits"
	"strconv"
)

//483. 最小好进制
//对于给定的整数 n, 如果n的k（k>=2）进制数的所有数位全为1，则称 k（k>=2）是 n 的一个好进制。
//
//以字符串的形式给出 n, 以字符串的形式返回 n 的最小好进制。
//
// 
//
//示例 1：
//
//输入："13"
//输出："3"
//解释：13 的 3 进制是 111。
//示例 2：
//
//输入："4681"
//输出："8"
//解释：4681 的 8 进制是 11111。
//示例 3：
//
//输入："1000000000000000000"
//输出："999999999999999999"
//解释：1000000000000000000 的 999999999999999999 进制是 11。
// 
//
//提示：
//
//n的取值范围是 [3, 10^18]。
//输入总是有效且没有前导 0。

//思路  数学
// m 范围  小于 Log_2_n
// k = [pow(n,1/m]
func smallestGoodBase(n string) string {
	N, _ := strconv.Atoi(n)
	mmax := bits.Len(uint(N)) - 1
	for m := mmax; m > -1; m-- {
		k := int(math.Pow(float64(N), 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if N == sum {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(N - 1)
}
