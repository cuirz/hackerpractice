package main

//1036. 逃离大迷宫
//在一个 10^6 x 10^6 的网格中，每个网格上方格的坐标为(x, y) 。
//
//现在从源方格source = [sx, sy]开始出发，意图赶往目标方格target = [tx, ty] 。数组 blocked 是封锁的方格列表，其中每个 blocked[i] = [xi, yi] 表示坐标为 (xi, yi) 的方格是禁止通行的。
//
//每次移动，都可以走到网格中在四个方向上相邻的方格，只要该方格 不 在给出的封锁列表blocked上。同时，不允许走出网格。
//
//只有在可以通过一系列的移动从源方格source 到达目标方格target 时才返回true。否则，返回 false。
//
//
//
//示例 1：
//
//输入：blocked = [[0,1],[1,0]], source = [0,0], target = [0,2]
//输出：false
//解释：
//从源方格无法到达目标方格，因为我们无法在网格中移动。
//无法向北或者向东移动是因为方格禁止通行。
//无法向南或者向西移动是因为不能走出网格。
//示例 2：
//
//输入：blocked = [], source = [0,0], target = [999999,999999]
//输出：true
//解释：
//因为没有方格被封锁，所以一定可以到达目标方格。
//
//
//提示：
//
//0 <= blocked.length <= 200
//blocked[i].length == 2
//0 <= xi, yi < 10^6
//source.length == target.length == 2
//0 <= sx, sy, tx, ty < 10^6
//source != target
//题目数据保证 source 和 target 不在封锁列表内

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	type pair struct {
		x, y int
	}
	dirs := []int{-1, 0, 1, 0, -1}
	const (
		block = iota
		valid
		found
		bound = 1e6
	)
	n := len(blocked)
	if n < 2 {
		return true
	}
	blockMap := map[pair]bool{}
	for _, b := range blocked {
		blockMap[pair{b[0], b[1]}] = true
	}
	bfs := func(s, t []int) int {
		sx, sy := s[0], s[1]
		tx, ty := t[0], t[1]
		limit := n * (n - 1) >> 1
		queue := []pair{{sx, sy}}
		visited := map[pair]bool{{sx, sy}: true}
		for len(queue) > 0 && limit > 0 {
			top := queue[0]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				x, y := top.x+dirs[i], top.y+dirs[i+1]
				p := pair{x, y}
				if x >= 0 && x < bound && y >= 0 && y < bound && !blockMap[p] && !visited[p] {
					if x == tx && y == ty {
						return found
					}
					limit--
					visited[p] = true
					queue = append(queue, p)
				}

			}
		}
		if limit > 0 {
			return block
		}
		return valid
	}
	result := bfs(source, target)
	return result == found || result == valid && bfs(target, source) != block

}
