package main

//723. 粉碎糖果
//这个问题是实现一个简单的消除算法。
//
//给定一个二维整数数组 board 代表糖果所在的方格，不同的正整数 board[i][j] 代表不同种类的糖果，如果 board[i][j] = 0 代表 (i, j) 这个位置是空的。给定的方格是玩家移动后的游戏状态，现在需要你根据以下规则粉碎糖果，使得整个方格处于稳定状态并最终输出。
//
//如果有三个及以上水平或者垂直相连的同种糖果，同一时间将它们粉碎，即将这些位置变成空的。
//在同时粉碎掉这些糖果之后，如果有一个空的位置上方还有糖果，那么上方的糖果就会下落直到碰到下方的糖果或者底部，这些糖果都是同时下落，也不会有新的糖果从顶部出现并落下来。
//通过前两步的操作，可能又会出现可以粉碎的糖果，请继续重复前面的操作。
//当不存在可以粉碎的糖果，也就是状态稳定之后，请输出最终的状态。
//你需要模拟上述规则并使整个方格达到稳定状态，并输出。

//思路 ：双指针
func candyCrush(board [][]int) [][]int {
	n := len(board)
	m := len(board[0])
	need := false

	tmp := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// row
			if i+2 < n && board[i][j] != 0 && abs(board[i][j]) == abs(board[i+1][j]) && abs(board[i][j]) == abs(board[i+2][j]) {
				tmp = -abs(board[i][j])
				board[i][j], board[i+1][j], board[i+2][j] = tmp, tmp, tmp
				need = true
			}
			// col
			if j+2 < m && board[i][j] != 0 && abs(board[i][j]) == abs(board[i][j+1]) && abs(board[i][j]) == abs(board[i][j+2]) {
				tmp = -abs(board[i][j])
				board[i][j], board[i][j+1], board[i][j+2] = tmp, tmp, tmp
				need = true
			}

		}
	}

	// 类似stack 一样的操作
	for j := 0; j < m; j++ {
		index := n - 1
		for i := n - 1; i > -1; i-- {
			if board[i][j] > 0 {
				board[index][j] = board[i][j]
				index--
			}
		}
		for i := index; i > -1; i-- {
			board[i][j] = 0
		}
	}
	if need {
		return candyCrush(board)
	}
	return board
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main() {
	array := [][]int{
		{110, 5, 112, 113, 114},
		{210, 211, 5, 213, 214},
		{310, 311, 3, 313, 314},
		{410, 411, 412, 5, 414},
		{5, 1, 512, 3, 3},
		{610, 4, 1, 613, 614},
		{710, 1, 2, 713, 714},
		{810, 1, 2, 1, 1},
		{1, 1, 2, 2, 2},
		{4, 1, 4, 4, 1014},
	}
	println(candyCrush(array))
}
