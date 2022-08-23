package main

import "math/bits"

//782. 变为棋盘
//一个n x n的二维网络board仅由0和1组成。每次移动，你能任意交换两列或是两行的位置。
//
//返回 将这个矩阵变为 “棋盘”所需的最小移动次数。如果不存在可行的变换，输出 -1。
//
//“棋盘” 是指任意一格的上下左右四个方向的值均与本身不同的矩阵。
//
//
//
//示例 1:
//
//
//
//输入: board = [[0,1,1,0],[0,1,1,0],[1,0,0,1],[1,0,0,1]]
//输出: 2
//解释:一种可行的变换方式如下，从左到右：
//第一次移动交换了第一列和第二列。
//第二次移动交换了第二行和第三行。
//示例 2:
//
//
//
//输入: board = [[0, 1], [1, 0]]
//输出: 0
//解释: 注意左上角的格值为0时也是合法的棋盘，也是合法的棋盘.
//示例 3:
//
//
//
//输入: board = [[1, 0], [1, 0]]
//输出: -1
//解释: 任意的变换都不能使这个输入变为合法的棋盘。
//
//
//提示：
//
//n == board.length
//n == board[i].length
//2 <= n <= 30
//board[i][j]将只包含0或1

func movesToChessboard(board [][]int) int {
	n := len(board)
	rowMask, colMask := 0, 0
	for i := 0; i < n; i++ {
		rowMask |= board[0][i] << i
		colMask |= board[i][0] << i
	}
	reverseRowMask := ((1<<n) - 1) ^ rowMask
	reverseColMask := ((1<<n) - 1) ^ colMask
	rowCnt, colCnt := 0, 0
	for i := 0; i < n; i++ {
		currRowMask, currColMask := 0, 0
		for j := 0; j < n; j++ {
			currRowMask |= board[i][j] << j
			currColMask |= board[j][i] << j
		}
		if currRowMask != rowMask && currRowMask != reverseRowMask ||
			currColMask != colMask && currColMask != reverseColMask {
			return -1
		}
		if currRowMask == rowMask {
			rowCnt++
		}
		if currColMask == colMask {
			colCnt++
		}
	}
	getMoves := func(mask uint, count, size int) int {
		ones := bits.OnesCount(mask)
		if size&1 > 0 {
			if abs(size-2*ones) != 1 || abs(size-2*count) != 1 {
				return -1
			}
			if ones == size>>1 {
				return size/2 - bits.OnesCount(mask&0xAAAAAAAA)
			} else {
				return (size+1)/2 - bits.OnesCount(mask&0x55555555)
			}
		} else {
			if ones != size>>1 || count != size>>1 {
				return -1
			}
			count0 := size/2 - bits.OnesCount(mask&0xAAAAAAAA)
			count1 := size/2 - bits.OnesCount(mask&0x55555555)
			return min(count0, count1)
		}

	}
	rowMoves := getMoves(uint(rowMask), rowCnt, n)
	colMoves := getMoves(uint(colMask), colCnt, n)
	if rowMoves == -1 || colMoves == -1 {
		return -1
	}
	return rowMoves + colMoves
}
