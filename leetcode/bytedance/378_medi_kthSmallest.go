package main

//378. 有序矩阵中第K小的元素
//给定一个n x n矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
//请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
//
//
//示例：
//
//matrix = [
//[ 1,  5,  9],
//[10, 11, 13],
//[12, 13, 15]
//],
//k = 8,
//
//返回 13。
//
//提示：
//你可以假设 k 的值永远是有效的，1 ≤ k ≤ n2。

//思路
// 二分法
// 当寻找的 mid值，左上角都会比它小，右下角都比它大。
// 寻找分割线

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l := matrix[0][0]
	r := matrix[n-1][n-1]
	var mid int
	check := func() bool {
		i := n - 1
		j := 0
		nums := 0
		for i >= 0 && j < n {
			if matrix[i][j] <= mid {
				nums += i + 1
				j++
			} else {
				i--
			}
		}
		return nums >= k
	}
	for l < r {
		mid = l + (r-l)>>1
		if check() {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	//println(kthSmallest([][]int{
	//	{1,5,9},
	//	{10,11,13},
	//	{12,13,15},
	//},8))
	println(kthSmallest([][]int{
		{1, 2},
		{1, 3},
	}, 2))
}
