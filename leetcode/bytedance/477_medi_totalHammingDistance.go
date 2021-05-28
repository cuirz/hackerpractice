package main

import "math/bits"

//477. 汉明距离总和
//两个整数的 汉明距离 指的是这两个数字的二进制数对应位不同的数量。
//
//计算一个数组中，任意两个数之间汉明距离的总和。
//
//示例:
//
//输入: 4, 14, 2
//
//输出: 6
//
//解释: 在二进制表示中，4表示为0100，14表示为1110，2表示为0010。（这样表示是为了体现后四位之间关系）
//所以答案为：
//HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.
//注意:
//
//数组中元素的范围为从 0到 10^9。
//数组的长度不超过 10^4。

func totalHammingDistance(nums []int) int {
	maxn := 0
	for _, v := range nums {
		if maxn < v {
			maxn = v
		}
	}
	L := bits.Len(uint(maxn))
	n := len(nums)
	var result int
	for i := 0; i <= L; i++ {
		count := 0
		for _, v := range nums {
			if v>>i&1 == 1 {
				count++
			}
		}
		result += count * (n - count)
	}
	return result
}

func main() {
	println(totalHammingDistance([]int{4, 14, 2}))
}
