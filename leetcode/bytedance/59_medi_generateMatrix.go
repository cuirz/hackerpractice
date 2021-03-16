package main

//59. 螺旋矩阵 II
//给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//示例 2：
//
//输入：n = 1
//输出：[[1]]
// 
//
//提示：
//
//1 <= n <= 20

func generateMatrix(n int) [][]int {
	r, d, l, u := n-1, n-1, 0, 0
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	index := 1
	for l <= r && u <= d {
		for col := l; col <= r; col++ {
			matrix[u][col] = index
			index++
		}
		for row := u+1; row <= d; row++ {
			matrix[row][r] = index
			index++
		}
		if l < r && u < d {
			for col := r-1; col > l; col-- {
				matrix[d][col] = index
				index++
			}
			for row := d; row > u; row-- {
				matrix[row][l] = index
				index++
			}
		}
		l++
		u++
		r--
		d--

	}
	return matrix
}
