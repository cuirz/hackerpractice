package main

//79. 单词搜索
//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
//示例:
//
//board =
//[
//['A','B','C','E'],
//['S','F','C','S'],
//['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false
//
//
//提示：
//
//board 和 word 中只包含大写和小写英文字母。
//1 <= board.length <= 200
//1 <= board[i].length <= 200
//1 <= word.length <= 10^3

//思路  dfs
func exist(board [][]byte, word string) bool {
	dir := []int{-1, 0, 1, 0, -1}
	n := len(board)
	m := len(board[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	isMatched := false
	var dfs func(i, j, index int)
	dfs = func(i, j, index int) {
		if i < 0 || j < 0 || i >= n || j >= m || visited[i][j] ||
			board[i][j] != word[index] || isMatched {
			return
		}
		if index == len(word)-1 {
			isMatched = true
			return
		}
		visited[i][j] = true
		for k := 0; k < 4; k++ {
			x := i + dir[k]
			y := j + dir[k+1]
			dfs(x, y, index+1)
		}
		visited[i][j] = false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !isMatched {
				dfs(i, j, 0)
			} else {
				return true
			}
		}
	}
	return isMatched
}
