package main

//矩阵中的最长递增路径
//输入: nums =
//[
//[9,9,4],
//[6,6,8],
//[2,1,1]
//]
//输出: 4
//解释: 最长递增路径为 [1, 2, 6, 9]。

func longestIncreasingPath(matrix [][]int) int {
	n := len(matrix)
	if n < 1 {
		return 0
	}
	m := len(matrix[0])

	count := make([][]int, n)
	for i := 0; i < n; i++ {
		count[i] = make([]int, m)
	}

	dir := []int{-1, 0, 1, 0, -1}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if count[i][j] != 0 {
			return count[i][j]
		}
		x, y := 0, 0
		for c := 0; c < 4; c++ {
			x = i + dir[c]
			y = j + dir[c+1]
			if x >= 0 && y >= 0 && x < n && y < m && matrix[x][y] > matrix[i][j] {
				count[i][j] = max(dfs(x, y), count[i][j])
			}
		}
		count[i][j]++
		return count[i][j]
	}
	var result int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result = max(result, dfs(i, j))
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
