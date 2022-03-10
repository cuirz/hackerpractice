package main

//LCP 13. 寻宝
//我们得到了一副藏宝图，藏宝图显示，在一个迷宫中存在着未被世人发现的宝藏。
//迷宫是一个二维矩阵，用一个字符串数组表示。它标识了唯一的入口（用 'S' 表示），和唯一的宝藏地点（用 'T' 表示）。但是，宝藏被一些隐蔽的机关保护了起来。在地图上有若干个机关点（用 'M' 表示），只有所有机关均被触发，才可以拿到宝藏。
//要保持机关的触发，需要把一个重石放在上面。迷宫中有若干个石堆（用 'O' 表示），每个石堆都有无限个足够触发机关的重石。但是由于石头太重，我们一次只能搬一个石头到指定地点。
//迷宫中同样有一些墙壁（用 '#' 表示），我们不能走入墙壁。剩余的都是可随意通行的点（用 '.' 表示）。石堆、机关、起点和终点（无论是否能拿到宝藏）也是可以通行的。
//我们每步可以选择向上/向下/向左/向右移动一格，并且不能移出迷宫。搬起石头和放下石头不算步数。那么，从起点开始，我们最少需要多少步才能最后拿到宝藏呢？如果无法拿到宝藏，返回 -1 。
//
//示例 1：
//输入： ["S#O", "M..", "M.T"]
//输出：16
//
//解释：最优路线为： S->O, cost = 4, 去搬石头 O->第二行的M, cost = 3, M机关触发 第二行的M->O, cost = 3, 我们需要继续回去 O 搬石头。 O->第三行的M, cost = 4, 此时所有机关均触发 第三行的M->T, cost = 2，去T点拿宝藏。 总步数为16。
//示例 2：
//
//输入： ["S#O", "M.#", "M.T"]
//输出：-1
//解释：我们无法搬到石头触发机关
//
//示例 3：
//输入： ["S#O", "M.T", "M.."]
//输出：17
//
//解释：注意终点也是可以通行的。
//限制：
//1 <= maze.length<= 100
//1 <= maze[i].length<= 100
//maze[i].length == maze[j].length
//S 和 T 有且只有一个
//0 <= M的数量 <= 16
//0 <= O的数量 <= 40，题目保证当迷宫中存在 M 时，一定存在至少一个 O 。

//思路 状态压缩动态规划

var (
	dir  = []int{-1, 0, 1, 0, -1}
	n, m int
)

func minimalSteps(maze []string) int {
	n = len(maze)
	m = len(maze[0])
	var sx, sy, tx, ty int
	var triggers, stones [][2]int

	for i := 0; i < n; i++ {
		for j, v := range maze[i] {
			switch v {
			case 'S':
				sx, sy = i, j
			case 'T':
				tx, ty = i, j
			case 'M':
				triggers = append(triggers, [2]int{i, j})
			case 'O':
				stones = append(stones, [2]int{i, j})
			}
		}
	}

	tCount, sCount := len(triggers), len(stones)

	startDist := bfs(sx, sy, maze)
	//没有机关时，直接到终点
	if tCount == 0 {
		return startDist[tx][ty]
	}

	// 机关 to 机关 最短路径计算 m:M
	dist := make([][]int, tCount)
	tPath := make([][][]int, tCount)
	for i := 0; i < tCount; i++ {
		// 扩展2个位置，一个是存放 起点to石头to机关 最短距，另一个是 机关到终点 最短距
		dist[i] = make([]int, tCount+2)
		for j := 0; j < tCount+2; j++ {
			dist[i][j] = -1
		}
		tPath[i] = bfs(triggers[i][0], triggers[i][1], maze)

		//机关到终点的最短距离
		dist[i][tCount+1] = tPath[i][tx][ty]
	}
	// 机关 to 石头 to 机关 ,计算最短路径 m:O:M
	for i := 0; i < tCount; i++ {
		minimum := -1
		// 先计算  起点 to 石头 to 机关
		// 每个机关必须遍历所有石头，才能判断最短路径
		for j := 0; j < sCount; j++ {
			ox := stones[j][0]
			oy := stones[j][1]
			if startDist[ox][oy] != -1 && tPath[i][ox][oy] != -1 {
				//minimum = minPositive(minimum, startDist[ox][oy]+tPath[i][ox][oy])
				if minimum == -1 || minimum > startDist[ox][oy]+tPath[i][ox][oy] {
					minimum = startDist[ox][oy] + tPath[i][ox][oy]
				}
			}
		}
		dist[i][tCount] = minimum
		// 机关 to 石头 to 机关
		for k := i + 1; k < tCount; k++ {
			minimum = -1
			for j := 0; j < sCount; j++ {
				ox := stones[j][0]
				oy := stones[j][1]
				if tPath[i][ox][oy] != -1 {
					//minimum = minPositive(minimum, tPath[i][ox][oy]+tPath[k][ox][oy])
					if minimum == -1 || minimum > tPath[i][ox][oy]+tPath[k][ox][oy] {
						minimum = tPath[i][ox][oy] + tPath[k][ox][oy]
					}
				}
			}
			dist[i][k] = minimum
			dist[k][i] = minimum
		}
	}

	// 无法到达
	for i := 0; i < tCount; i++ {
		if dist[i][tCount] == -1 || dist[i][tCount+1] == -1 {
			return -1
		}
	}

	// 动态规划
	size := 1 << tCount
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, tCount)
		for j := 0; j < tCount; j++ {
			dp[i][j] = -1
		}
	}

	for i := 0; i < tCount; i++ {
		dp[1<<i][i] = dist[i][tCount]
	}

	// 机关激活状态 state表达，每位代表一个机关. 遍历所有可能的状态
	for state := 1; state < (1 << tCount); state++ {
		for i := 0; i < tCount; i++ {
			// 把自己过滤
			if state&(1<<i) != 0 {
				for j := 0; j < tCount; j++ {
					// 已经激活的过滤
					if state&(1<<j) == 0 {
						nextState := state | (1 << j)
						//dp[nextState][j] = minPositive(dp[nextState][j], dp[state][i]+dist[i][j])
						if dp[nextState][j] == -1 || dp[nextState][j] > dp[state][i]+dist[i][j] {
							dp[nextState][j] = dp[state][i] + dist[i][j]
						}
					}
				}
			}
		}
	}

	// 所有已经被激活状态
	allActivatedState := size - 1
	result := -1
	for i := 0; i < tCount; i++ {
		//result = minPositive(result, dp[allActivatedState][i]+dist[i][tCount+1])
		if result == -1 || result > dp[allActivatedState][i]+dist[i][tCount+1] {
			result = dp[allActivatedState][i] + dist[i][tCount+1]
		}
	}
	return result
}

func bfs(x, y int, maze []string) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = -1
		}
	}
	result[x][y] = 0
	queue := make([][]int, 0)
	queue = append(queue, []int{x, y})
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			nx := first[0] + dir[i]
			ny := first[1] + dir[i+1]
			if nx >= 0 && ny >= 0 && nx < n && ny < m && maze[nx][ny] != '#' && result[nx][ny] == -1 {
				result[nx][ny] = result[first[0]][first[1]] + 1
				queue = append(queue, []int{nx, ny})
			}
		}
	}
	return result
}

func minPositive(x, y int) int {
	println(x, y)
	if x*y > 0 {
		return min(x, y)
	}
	if x < 0 {
		return y
	}
	return x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	//["...O.",".S#M.","..#T.","....."]  5
	minimalSteps([]string{"S#O", "M..", "M.T"})
}
