package main

//847. 访问所有节点的最短路径
//存在一个由 n 个节点组成的无向连通图，图中的节点按从 0 到 n - 1 编号。
//
//给你一个数组 graph 表示这个图。其中，graph[i] 是一个列表，由所有与节点 i 直接相连的节点组成。
//
//返回能够访问所有节点的最短路径的长度。你可以在任一节点开始和停止，也可以多次重访节点，并且可以重用边。
//
//
//
//示例 1：
//
//
//输入：graph = [[1,2,3],[0],[0],[0]]
//输出：4
//解释：一种可能的路径为 [1,0,2,0,3]
//示例 2：
//
//
//
//输入：graph = [[1],[0,2,4],[1,3,4],[2],[1,2]]
//输出：4
//解释：一种可能的路径为 [0,1,4,2,3]
//
//
//提示：
//
//n == graph.length
//1 <= n <= 12
//0 <= graph[i].length <n
//graph[i] 不包含 i
//如果 graph[a] 包含 b ，那么 graph[b] 也包含 a
//输入的图总是连通图

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	seen := make([][]bool, n)
	type tuple struct{ u, mask, dist int }
	queue := make([]tuple, 0)
	for i := 0; i < n; i++ {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		queue = append(queue, tuple{i, 1 << i, 0})
	}
	for {
		top := queue[0]
		queue = queue[1:]
		if top.mask == (1<<n)-1 {
			return top.dist
		}
		for _, v := range graph[top.u] {
			mask := top.mask | 1<<v
			if !seen[v][mask] {
				queue = append(queue, tuple{v, mask, top.dist + 1})
				seen[v][mask] = true
			}
		}
	}
}
