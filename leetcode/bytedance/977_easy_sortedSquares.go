package main

//977. 有序数组的平方
//给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
//
//
//
//示例 1：
//
//输入：[-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//示例 2：
//
//输入：[-7,-3,2,3,11]
//输出：[4,9,9,49,121]
//
//
//提示：
//
//1 <= A.length <= 10000
//-10000 <= A[i] <= 10000
//A已按非递减顺序排序。

func sortedSquares(A []int) []int {
	n := len(A)
	result := make([]int, n)
	h, e := 0, n-1
	for seek := n - 1; seek >= 0; seek-- {
		x := A[h] * A[h]
		y := A[e] * A[e]
		if x > y {
			result[seek] = x
			h++
		} else {
			result[seek] = y
			e--
		}
	}
	return result
}
