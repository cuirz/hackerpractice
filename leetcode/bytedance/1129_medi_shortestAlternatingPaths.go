package main

//1129. 颜色交替的最短路径
//在一个有向图中，节点分别标记为 0, 1, ..., n-1。图中每条边为红色或者蓝色，且存在自环或平行边。
//
//red_edges 中的每一个 [i, j] 对表示从节点 i 到节点 j 的红色有向边。类似地，blue_edges 中的每一个 [i, j] 对表示从节点 i 到节点 j 的蓝色有向边。
//
//返回长度为 n 的数组 answer，其中 answer[X] 是从节点 0 到节点 X 的红色边和蓝色边交替出现的最短路径的长度。如果不存在这样的路径，那么 answer[x] = -1。
//
//
//
//示例 1：
//
//输入：n = 3, red_edges = [[0,1],[1,2]], blue_edges = []
//输出：[0,1,-1]
//示例 2：
//
//输入：n = 3, red_edges = [[0,1]], blue_edges = [[2,1]]
//输出：[0,1,-1]
//示例 3：
//
//输入：n = 3, red_edges = [[1,0]], blue_edges = [[2,1]]
//输出：[0,-1,-1]
//示例 4：
//
//输入：n = 3, red_edges = [[0,1]], blue_edges = [[1,2]]
//输出：[0,1,2]
//示例 5：
//
//输入：n = 3, red_edges = [[0,1],[0,2]], blue_edges = [[1,0]]
//输出：[0,1,1]
//
//
//提示：
//
//1 <= n <= 100
//red_edges.length <= 400
//blue_edges.length <= 400
//red_edges[i].length == blue_edges[i].length == 2
//0 <= red_edges[i][j], blue_edges[i][j] < n
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	type pair struct {
		val   int
		color int
	}
	const red, blue = 0, 1
	graph := make([][]pair, n)
	for _, v := range redEdges {
		graph[v[0]] = append(graph[v[0]], pair{v[1], red})
	}
	for _, v := range blueEdges {
		graph[v[0]] = append(graph[v[0]], pair{v[1], blue})
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	visited := make([][2]bool, n)
	visited[0] = [2]bool{true, true}
	queue := []pair{{0, red}, {0, blue}}
	for lv := 0; len(queue) > 0; lv++ {
		tmp := queue
		queue = nil
		for _, p := range tmp {
			x := p.val
			if dis[x] < 0 {
				dis[x] = lv
			}
			for _, to := range graph[x] {
				if to.color != p.color && !visited[to.val][to.color] {
					visited[to.val][to.color] = true
					queue = append(queue, to)
				}
			}
		}
	}
	return dis

}
