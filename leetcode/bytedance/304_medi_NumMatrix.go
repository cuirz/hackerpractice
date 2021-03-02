package main

//304. 二维区域和检索 - 矩阵不可变
//给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
//给定 matrix = [
//[3, 0, 1, 4, 2],
//[5, 6, 3, 2, 1],
//[1, 2, 0, 1, 5],
//[4, 1, 0, 1, 7],
//[1, 0, 3, 0, 5]
//]
//
//sumRegion(2, 1, 4, 3) -> 8
//sumRegion(1, 1, 2, 2) -> 11
//sumRegion(1, 2, 2, 4) -> 12
//提示：
//
//你可以假设矩阵不可变。
//会多次调用 sumRegion 方法。
//你可以假设 row1 ≤ row2 且 col1 ≤ col2 。

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	if n == 0 {
		return NumMatrix{}
	}
	m := len(matrix[0])
	sums := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sums[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			sums[i][j] = matrix[i-1][j-1] + sums[i][j-1] + sums[i-1][j] - sums[i-1][j-1]
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row2+1][col1] - this.sum[row1][col2+1] + this.sum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
