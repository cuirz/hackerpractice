package main

//面试题29. 顺时针打印矩阵
//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
//
//
//示例 1：
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//示例 2：
//
//输入：matrix =[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n < 1 {
		return []int{}
	}
	m := len(matrix[0])
	result := make([]int, n*m)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	index := 0
	block := 0
	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var drive func(i, j, d int)
	drive = func(i, j, d int) {
		if d > 3 {
			d = 0
		}
		block = 0
		visited[i][j] = true
		result[index] = matrix[i][j]
		index++
		for block < 2 {
			x := i + dir[d][1]
			y := j + dir[d][0]
			if x >= 0 && y >= 0 && x < n && y < m && !visited[x][y] {
				drive(x, y, d)
				return
			}
			block++
			d++
			if d > 3 {
				d = 0
			}
		}
	}
	drive(0, 0, 0)
	return result
}

func main() {
	println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))

}
