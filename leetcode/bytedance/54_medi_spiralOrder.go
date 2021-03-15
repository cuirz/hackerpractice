package main

//54. 螺旋矩阵
//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//提示：
//
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 10
//-100 <= matrix[i][j] <= 100

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	result := make([]int, m*n)
	index := 0
	r, d, l, u := n-1, m-1, 0, 0
	for l <= r && u <= d {
		for col := l; col <= r; col++ {
			result[index] = matrix[u][col]
			index++
		}
		for row := u + 1; row <= d; row++ {
			result[index] = matrix[row][r]
			index++
		}
		if l < r && u < d {
			for col := r - 1; col > l; col-- {
				result[index] = matrix[d][col]
				index++
			}
			for row := d; row > u; row-- {
				result[index] = matrix[row][l]
				index++
			}
		}

		l++
		u++
		r--
		d--
	}
	return result
}
