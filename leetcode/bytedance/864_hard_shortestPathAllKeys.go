package main

import "unicode"

//864. 获取所有钥匙的最短路径
//给定一个二维网格 grid ，其中：
//
//'.' 代表一个空房间
//'#' 代表一堵
//'@' 是起点
//小写字母代表钥匙
//大写字母代表锁
//我们从起点开始出发，一次移动是指向四个基本方向之一行走一个单位空间。我们不能在网格外面行走，也无法穿过一堵墙。如果途经一个钥匙，我们就把它捡起来。除非我们手里有对应的钥匙，否则无法通过锁。
//
//假设 k 为 钥匙/锁 的个数，且满足 1 <= k <= 6，字母表中的前 k 个字母在网格中都有自己对应的一个小写和一个大写字母。换言之，每个锁有唯一对应的钥匙，每个钥匙也有唯一对应的锁。另外，代表钥匙和锁的字母互为大小写并按字母顺序排列。
//
//返回获取所有钥匙所需要的移动的最少次数。如果无法获取所有钥匙，返回 -1 。
//
//
//
//示例 1：
//
//
//
//输入：grid = ["@.a.#","###.#","b.A.B"]
//输出：8
//解释：目标是获得所有钥匙，而不是打开所有锁。
//示例 2：
//
//
//
//输入：grid = ["@..aA","..B#.","....b"]
//输出：6
//示例 3:
//
//
//输入: grid = ["@Aa"]
//输出: -1
//
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 30
//grid[i][j] 只含有 '.', '#', '@', 'a'-'f' 以及 'A'-'F'
//钥匙的数目范围是 [1, 6]
//每个钥匙都对应一个 不同 的字母
//每个钥匙正好打开一个对应的锁
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	sx, sy := 0, 0
	keyToIdx := map[rune]int{}
	for i, row := range grid {
		for j, c := range row {
			if c == '@' {
				sx, sy = i, j
			} else if unicode.IsLower(c) {
				if _, ok := keyToIdx[c]; !ok {
					keyToIdx[c] = len(keyToIdx)
				}
			}
		}
	}

	dist := make([][][]int, m)
	for i := range dist {
		dist[i] = make([][]int, n)
		for j := range dist[i] {
			dist[i][j] = make([]int, 1<<len(keyToIdx))
			for k := range dist[i][j] {
				dist[i][j][k] = -1
			}
		}
	}
	dist[sx][sy][0] = 0
	q := [][3]int{{sx, sy, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y, mask := p[0], p[1], p[2]
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] != '#' {
				c := rune(grid[nx][ny])
				if c == '.' || c == '@' {
					if dist[nx][ny][mask] == -1 {
						dist[nx][ny][mask] = dist[x][y][mask] + 1
						q = append(q, [3]int{nx, ny, mask})
					}
				} else if unicode.IsLower(c) {
					t := mask | 1<<keyToIdx[c]
					if dist[nx][ny][t] == -1 {
						dist[nx][ny][t] = dist[x][y][mask] + 1
						if t == 1<<len(keyToIdx)-1 {
							return dist[nx][ny][t]
						}
						q = append(q, [3]int{nx, ny, t})
					}
				} else {
					idx := keyToIdx[unicode.ToLower(c)]
					if mask>>idx&1 > 0 && dist[nx][ny][mask] == -1 {
						dist[nx][ny][mask] = dist[x][y][mask] + 1
						q = append(q, [3]int{nx, ny, mask})
					}
				}
			}
		}
	}
	return -1
}
