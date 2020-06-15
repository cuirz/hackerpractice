package main

//1314. 矩阵区域和
//给你一个 m * n 的矩阵 mat 和一个整数 K ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素 mat[r][c] 的和：
//
//i - K <= r <= i + K, j - K <= c <= j + K
//(r, c) 在矩阵内。
//
//
//示例 1：
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]], K = 1
//输出：[[12,21,16],[27,45,33],[24,39,28]]
//示例 2：
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]], K = 2
//输出：[[45,45,45],[45,45,45],[45,45,45]]

//方案  前缀和

func matrixBlockSum(mat [][]int, K int) [][]int {
	n := len(mat)
	m := len(mat[0])
	sums := make([][]int, n+1)

	sums[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		sums[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			sums[i][j] = mat[i-1][j-1] + sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1]
		}
	}
	lx, ly, rx, ry := 0, 0, 0, 0

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			lx = max(i-K, 0)
			ly = max(j-K, 0)
			rx = min(i+K, n-1)
			ry = min(j+K, m-1)
			result[i][j] = sums[rx+1][ry+1] - sums[lx][ry+1] - sums[rx+1][ly] + sums[lx][ly]
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
