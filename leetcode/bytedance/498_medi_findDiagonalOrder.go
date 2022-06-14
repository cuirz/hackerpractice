package main

//498. 对角线遍历
//给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
//
//
//
//示例 1：
//
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,4,7,5,3,6,8,9]
//示例 2：
//
//输入：mat = [[1,2],[3,4]]
//输出：[1,2,3,4]
//
//提示：
//
//m == mat.length
//n == mat[i].length
//1 <= m, n <= 10^4
//1 <= m * n <= 10^4
//-10^5 <= mat[i][j] <= 10^5

func findDiagonalOrder(mat [][]int) []int {
	n := len(mat)
	m := len(mat[0])
	l := m + n - 1
	var result []int
	for k := 0; k < l; k++ {
		if k%2 == 0 {
			x := max(0, k-n+1)
			y := k - x
			for x < m && y >= 0 {
				result = append(result, mat[y][x])
				x++
				y--
			}
		} else {
			y := max(0, k-m+1)
			x := k - y
			for x >= 0 && y < n {
				result = append(result, mat[y][x])
				x--
				y++
			}

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
func main() {
	println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
