package main

//454. 四数相加 II
//给定四个包含整数的数组列表A , B , C , D ,计算有多少个元组 (i, j, k, l)，使得A[i] + B[j] + C[k] + D[l] = 0。
//
//为了使问题简单化，所有的 A, B, C, D 具有相同的长度N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过231 - 1 。
//
//例如:
//
//输入:
//A = [ 1, 2]
//B = [-2,-1]
//C = [-1, 2]
//D = [ 0, 2]
//
//输出:
//2
//
//解释:
//两个元组如下:
//1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
//2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0

//思路 分组哈希
func fourSumCount(A []int, B []int, C []int, D []int) int {
	count := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			count[a+b]++
		}
	}
	var result int
	for _, c := range C {
		for _, d := range D {
			result += count[-c-d]
		}
	}
	return result
}
