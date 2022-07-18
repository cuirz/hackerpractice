package main

//749. 隔离病毒
//病毒扩散得很快，现在你的任务是尽可能地通过安装防火墙来隔离病毒。
//
//假设世界由m x n的二维矩阵isInfected组成，isInfected[i][j] == 0表示该区域未感染病毒，而 isInfected[i][j] == 1表示该区域已感染病毒。可以在任意 2 个相邻单元之间的共享边界上安装一个防火墙（并且只有一个防火墙）。
//
//每天晚上，病毒会从被感染区域向相邻未感染区域扩散，除非被防火墙隔离。现由于资源有限，每天你只能安装一系列防火墙来隔离其中一个被病毒感染的区域（一个区域或连续的一片区域），且该感染区域对未感染区域的威胁最大且 保证唯一。
//
//你需要努力使得最后有部分区域不被病毒感染，如果可以成功，那么返回需要使用的防火墙个数; 如果无法实现，则返回在世界被病毒全部感染时已安装的防火墙个数。
//
//
//
//示例 1：
//
//
//
//输入: isInfected = [[0,1,0,0,0,0,0,1],[0,1,0,0,0,0,0,1],[0,0,0,0,0,0,0,1],[0,0,0,0,0,0,0,0]]
//输出: 10
//解释:一共有两块被病毒感染的区域。
//在第一天，添加 5 墙隔离病毒区域的左侧。病毒传播后的状态是:
//
//第二天，在右侧添加 5 个墙来隔离病毒区域。此时病毒已经被完全控制住了。
//
//示例 2：
//
//
//
//输入: isInfected = [[1,1,1],[1,0,1],[1,1,1]]
//输出: 4
//解释: 虽然只保存了一个小区域，但却有四面墙。
//注意，防火墙只建立在两个不同区域的共享边界上。
//示例3:
//
//输入: isInfected = [[1,1,1,0,0,0,0,0,0],[1,0,1,0,1,1,1,1,1],[1,1,1,0,0,0,0,0,0]]
//输出: 13
//解释: 在隔离右边感染区域后，隔离左边病毒区域只需要 2 个防火墙。
//
//
//提示:
//
//m ==isInfected.length
//n ==isInfected[i].length
//1 <= m, n <= 50
//isInfected[i][j]is either0or1
//在整个描述的过程中，总有一个相邻的病毒区域，它将在下一轮 严格地感染更多未受污染的方块

// 广度优先
func containVirus(isInfected [][]int) int {
	m, n := len(isInfected), len(isInfected[0])
	type pair struct {
		x, y int
	}
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var ans int
	for {
		var neighbors []map[pair]struct{}
		var firwalls []int
		for i, row := range isInfected {
			for j, infected := range row {
				if infected != 1 {
					continue
				}
				q := []pair{{i, j}}
				neighbor := map[pair]struct{}{}
				firwall, idx := 0, len(neighbors)+1
				row[j] = -idx
				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					for _, d := range dirs {
						if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n {
							if isInfected[x][y] == 1 {
								q = append(q, pair{x, y})
								isInfected[x][y] = -idx
							} else if isInfected[x][y] == 0 {
								firwall++
								neighbor[pair{x, y}] = struct{}{}
							}
						}
					}
				}
				neighbors = append(neighbors, neighbor)
				firwalls = append(firwalls, firwall)
			}
		}
		if len(neighbors) == 0 {
			return ans
		}
		idx := 0
		for i := 1; i < len(neighbors); i++ {
			if len(neighbors[i]) > len(neighbors[idx]) {
				idx = i
			}
		}
		ans += firwalls[idx]
		for _, row := range isInfected {
			for j, v := range row {
				if v < 0 {
					if v != -idx-1 {
						row[j] = 1

					} else {
						row[j] = 2
					}
				}
			}
		}
		for i, neighbor := range neighbors {
			if i != idx {
				for p := range neighbor {
					isInfected[p.x][p.y] = 1
				}
			}
		}
		if len(neighbors) == 1 {
			return ans
		}
	}
	return ans
}
