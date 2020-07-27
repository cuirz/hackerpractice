package main

//461. 汉明距离
//两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
//
//给出两个整数 x 和 y，计算它们之间的汉明距离。
//
//注意：
//0 ≤ x, y < 231.

//思路  异或

//思路 布赖恩·克尼根算法
//x  		= 		1 0 0 0 1 0 0 0
//x-1		=  		1 0 0 0 0 1 1 1
//x & (x-1) =		1 0 0 0 0 0 0 0
// 一步操作越过了 x 末尾 1之后的 0,减少了移位步骤
func hammingDistanceBest(x int, y int) int {
	z := x ^ y
	count := 0
	for z > 0 {
		count++
		z = z & (z - 1)
	}
	return count

}

func hammingDistance(x int, y int) int {
	z := x ^ y
	count := 0
	for z > 0 {
		if z&0x1 == 1 {
			count++
		}
		z >>= 1
	}
	return count

}
