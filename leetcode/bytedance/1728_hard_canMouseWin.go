package main

//1728. 猫和老鼠 II
//一只猫和一只老鼠在玩一个叫做猫和老鼠的游戏。
//
//它们所处的环境设定是一个rows x cols的方格 grid，其中每个格子可能是一堵墙、一块地板、一位玩家（猫或者老鼠）或者食物。
//
//玩家由字符'C'（代表猫）和'M'（代表老鼠）表示。
//地板由字符'.'表示，玩家可以通过这个格子。
//墙用字符'#'表示，玩家不能通过这个格子。
//食物用字符'F'表示，玩家可以通过这个格子。
//字符'C'，'M'和'F'在grid中都只会出现一次。
//猫和老鼠按照如下规则移动：
//
//老鼠 先移动，然后两名玩家轮流移动。
//每一次操作时，猫和老鼠可以跳到上下左右四个方向之一的格子，他们不能跳过墙也不能跳出grid。
//catJump 和mouseJump是猫和老鼠分别跳一次能到达的最远距离，它们也可以跳小于最大距离的长度。
//它们可以停留在原地。
//老鼠可以跳跃过猫的位置。
//游戏有 4 种方式会结束：
//
//如果猫跟老鼠处在相同的位置，那么猫获胜。
//如果猫先到达食物，那么猫获胜。
//如果老鼠先到达食物，那么老鼠获胜。
//如果老鼠不能在 1000 次操作以内到达食物，那么猫获胜。
//给你rows x cols的矩阵grid和两个整数catJump和mouseJump，双方都采取最优策略，如果老鼠获胜，那么请你返回true，否则返回 false。
//
//
//
//示例 1：
//
//
//
//输入：grid = ["####F","#C...","M...."], catJump = 1, mouseJump = 2
//输出：true
//解释：猫无法抓到老鼠，也没法比老鼠先到达食物。
//示例 2：
//
//
//
//输入：grid = ["M.C...F"], catJump = 1, mouseJump = 4
//输出：true
//示例 3：
//
//输入：grid = ["M.C...F"], catJump = 1, mouseJump = 3
//输出：false
//示例 4：
//
//输入：grid = ["C...#","...#F","....#","M...."], catJump = 2, mouseJump = 5
//输出：false
//示例 5：
//
//输入：grid = [".M...","..#..","#..#.","C#.#.","...#F"], catJump = 3, mouseJump = 1
//输出：true
//
//
//提示：
//
//rows == grid.length
//cols = grid[i].length
//1 <= rows, cols <= 8
//grid[i][j] 只包含字符'C'，'M'，'F'，'.'和'#'。
//grid中只包含一个'C'，'M'和'F'。
//1 <= catJump, mouseJump <= 8

//913 猫和老鼠 I 关联
// 方法I dp
// 方法II 拓扑方法

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	const (
		TURN_MOUSE = iota
		TURN_CAT
	)
	n, m := len(grid), len(grid[0])
	maxR := n * m
	// mx,my,cx,cy,turn,step
	var dp [8][8][8][8][2][70]bool
	dir := []int{-1, 0, 1, 0, -1}
	var _mx, _my, _cx, _cy, fx, fy int

	var dfs func(mx, my, cx, cy, turn, step int) bool
	dfs = func(mx, my, cx, cy, turn, step int) bool {
		if step == maxR {
			return false
		}
		if turn == TURN_CAT && mx == fx && my == fy {
			return false
		} else if turn == TURN_CAT && cx == mx && cy == my {
			return true
		} else if turn == TURN_MOUSE && cx == fx && cy == fy {
			return false
		} else if turn == TURN_MOUSE && mx == cx && my == cy {
			return false
		}
		if dp[mx][my][cx][cy][turn][step] {
			return dp[mx][my][cx][cy][turn][step]
		}
		var isWin bool
		if turn == TURN_MOUSE {
			if !dfs(mx, my, cx, cy, TURN_CAT, step) {
				isWin = true
			} else {
				for i := 0; i < 4 && !isWin; i++ {
					for dis := 1; dis <= mouseJump && !isWin; dis++ {
						x, y := mx+dis*dir[i], my+dis*dir[i+1]
						if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == '#' {
							break
						}
						if !dfs(x, y, cx, cy, TURN_CAT, step) {
							isWin = true
						}
					}
				}
			}
		} else {
			if !dfs(mx, my, cx, cy, TURN_MOUSE, step+1) {
				isWin = true
			}
			for i := 0; i < 4 && !isWin; i++ {
				for dis := 1; dis <= catJump && !isWin; dis++ {
					x, y := cx+dis*dir[i], cy+dis*dir[i+1]
					if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == '#' {
						break
					}
					if !dfs(mx, my, x, y, TURN_MOUSE, step+1) {
						isWin = true
					}
				}
			}
		}

		dp[mx][my][cx][cy][turn][step] = isWin
		return isWin
	}
	for x, rows := range grid {
		for y, block := range rows {
			switch block {
			case 'M':
				_mx, _my = x, y
			case 'C':
				_cx, _cy = x, y
			case 'F':
				fx, fy = x, y
			}
		}
	}
	return dfs(_mx, _my, _cx, _cy, 0, 1)

}
