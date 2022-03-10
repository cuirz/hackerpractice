package main

import "strings"

//51. N 皇后
//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//给定一个整数 n，返回所有不同的n皇后问题的解决方案。
//
//每一种解法包含一个明确的n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//示例：
//
//输入：4
//输出：[
//[".Q..",  // 解法 1
//"...Q",
//"Q...",
//"..Q."],
//
//["..Q.",  // 解法 2
//"Q...",
//"...Q",
//".Q.."]
//]
//解释: 4 皇后问题存在两个不同的解法。
//
//
//提示：
//
//皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	queue := make([]int, 0)

	finish := func() {
		res := make([]string, 0)
		line := &strings.Builder{}
		for i := 0; i < n; i++ {
			line.Reset()
			for j := 0; j < n; j++ {
				if j == queue[i] {
					line.WriteString("Q")
				} else {
					line.WriteString(".")
				}
			}
			res = append(res, line.String())
		}
		result = append(result, res)
	}
	check := func(i, j int) bool {
		for k := 0; k < len(queue); k++ {
			if queue[k] == j || i-k == abs(j-queue[k]) {
				return false
			}
		}
		return true
	}

	var recursive func(i int)
	recursive = func(i int) {
		if i == n {
			//生成 result
			finish()
			queue = queue[:len(queue)-1]
			return
		}

		for j := 0; j < n; j++ {
			if check(i, j) {
				queue = append(queue, j)
				recursive(i + 1)
			}
		}
		queue = queue[:len(queue)-1]
	}
	for i := 0; i < n; i++ {
		queue = append(queue, i)
		recursive(1)
	}
	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	solveNQueens(4)
}
