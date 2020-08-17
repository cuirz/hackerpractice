package main

//130. 被围绕的区域
//给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
//
//找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
//
//示例:
//
//X X X X
//X O O X
//X X O X
//X O X X
//运行你的函数后，矩阵变为：
//
//X X X X
//X X X X
//X X X X
//X O X X
//解释:
//
//被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

//思路  dfs
// 标记状态  先标记 边缘'o'相连接的'o'，最后剩下的就是被关起来的

func solve(board [][]byte) {
	dir := []int{-1, 0, 1, 0, -1}
	n := len(board)
	if n < 1{
		return
	}
	m := len(board[0])
	var search func(i, j int)
	search = func(i, j int) {
		if i < 0 || j < 0 || i >= n || j >= m || board[i][j] != 'O' {
			return
		}
		board[i][j] = 0
		for k := 0; k < 4; k++ {
			x := i + dir[k]
			y := j + dir[k+1]
			search(x, y)
		}

	}
	for j := 0; j < m; j++ {
		search(0, j)
		search(n-1, j)
	}
	for i := 0; i < n; i++ {
		search(i, 0)
		search(i, m-1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 0{
				board[i][j] = 'O'
			}else if board[i][j] == 'O'{
				board[i][j] = 'X'
			}
		}
	}

}
