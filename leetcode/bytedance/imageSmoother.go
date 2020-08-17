package main

//包含整数的二维矩阵 M 表示一个图片的灰度。你需要设计一个平滑器来让每一个单元的灰度成为平均灰度 (向下舍入) ，平均灰度的计算是周围的8个单元和它本身的值求平均，如果周围的单元格不足八个，则尽可能多的利用它们。

func imageSmoother(M [][]int) [][]int {
	dir := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	n := len(M)
	m := len(M[0])
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sum := 0
			count := 0
			for k := 0; k < len(dir); k++ {
				x := i + dir[k][0]
				y := j + dir[k][1]
				if x >= 0 && y >= 0 && x < n && y < m {
					sum +=M[x][y]
					count ++
				}
			}
			result[i][j] = sum / count
		}
	}
	return result
}

func main(){
	println(imageSmoother([][]int{{1,1},{1,2}}))
}
