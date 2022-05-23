package main

import "sort"

//675. 为高尔夫比赛砍树
//你被请来给一个要举办高尔夫比赛的树林砍树。树林由一个m x n 的矩阵表示， 在这个矩阵中：
//
//0 表示障碍，无法触碰
//1表示地面，可以行走
//比 1 大的数表示有树的单元格，可以行走，数值表示树的高度
//每一步，你都可以向上、下、左、右四个方向之一移动一个单位，如果你站的地方有一棵树，那么你可以决定是否要砍倒它。
//
//你需要按照树的高度从低向高砍掉所有的树，每砍过一颗树，该单元格的值变为 1（即变为地面）。
//
//你将从 (0, 0) 点开始工作，返回你砍完所有树需要走的最小步数。 如果你无法砍完所有的树，返回 -1 。
//
//可以保证的是，没有两棵树的高度是相同的，并且你至少需要砍倒一棵树。
//
//
//
//示例 1：
//
//
//输入：forest = [[1,2,3],[0,0,4],[7,6,5]]
//输出：6
//解释：沿着上面的路径，你可以用 6 步，按从最矮到最高的顺序砍掉这些树。
//示例 2：
//
//
//输入：forest = [[1,2,3],[0,0,0],[7,6,5]]
//输出：-1
//解释：由于中间一行被障碍阻塞，无法访问最下面一行中的树。
//示例 3：
//
//输入：forest = [[2,3,4],[0,0,5],[8,7,6]]
//输出：6
//解释：可以按与示例 1 相同的路径来砍掉所有的树。
//(0,0) 位置的树，可以直接砍去，不用算步数。
//
//
//提示：
//
//m == forest.length
//n == forest[i].length
//1 <= m, n <= 50
//0 <= forest[i][j] <= 10^9

// 思路 先对树排序，然后累加树之间最短路径
func cutOffTree(forest [][]int) int {
	m := len(forest)
	n := len(forest[0])
	type pairs struct {
		h, x, y int
	}
	var trees []pairs
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				trees = append(trees, pairs{forest[i][j], i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool {
		return trees[i].h < trees[j].h
	})
	dirs := []int{-1, 0, 1, 0, -1}
	bfs := func(sx, sy, tx, ty int) int {
		visited := make([][]bool, m)
		for i := 0; i < m; i++ {
			visited[i] = make([]bool, n)
		}
		visited[sx][sy] = true
		queue := []pairs{{0, sx, sy}}
		for len(queue) > 0 {
			top := queue[0]
			queue = queue[1:]
			if top.x == tx && top.y == ty {
				return top.h
			}
			for i := 0; i < 4; i++ {
				x, y, dis := top.x+dirs[i], top.y+dirs[i+1], top.h
				if x >= 0 && y >= 0 && x < m && y < n && !visited[x][y] && forest[x][y] > 0 {
					visited[x][y] = true
					queue = append(queue, pairs{dis + 1, x, y})
				}
			}
		}
		return -1
	}
	prex, prey := 0, 0
	var result int
	for _, v := range trees {
		dis := bfs(prex, prey, v.x, v.y)
		if dis < 0 {
			return -1
		}
		result += dis
		prex, prey = v.x, v.y
	}
	return result
}

func main() {
	println(cutOffTree([][]int{{54581641, 64080174, 24346381, 69107959}, {86374198, 61363882, 68783324, 79706116}, {668150, 92178815, 89819108, 94701471}, {83920491, 22724204, 46281641, 47531096}, {89078499, 18904913, 25462145, 60813308}}))
}
