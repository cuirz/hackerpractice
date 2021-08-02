package main

import "math"

//743. 网络延迟时间
//有 n 个网络节点，标记为 1 到 n。
//
//给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。
//
//现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
// 
//
//示例 1：
//
//
//
//输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
//输出：2
//示例 2：
//
//输入：times = [[1,2,1]], n = 2, k = 1
//输出：1
//示例 3：
//
//输入：times = [[1,2,1]], n = 2, k = 2
//输出：-1
// 
//
//提示：
//
//1 <= k <= n <= 100
//1 <= times.length <= 6000
//times[i].length == 3
//1 <= ui, vi <= n
//ui != vi
//0 <= wi <= 100
//所有 (ui, vi) 对都 互不相同（即，不含重复边）

func networkDelayTime(times [][]int, n int, k int) int {
	const INF = math.MaxInt64 / 2
	g := make([][]int, n)

	for x := range g {
		g[x] = make([]int, n)
		for y := range g {
			g[x][y] = INF
		}
	}
	for _, v := range times {
		x, y := v[0]-1, v[1]-1
		g[x][y] = v[2]
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[k-1] = 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		for y, v := range visited {
			if !v && (x == -1 || dist[y] < dist[x]) {
				x = y
			}
		}
		visited[x] = true
		for y, time := range g[x] {
			dist[y] = min(dist[y], dist[x]+time)
		}
	}

	var result int
	for i, v := range dist {
		if dist[i] == INF {
			return -1
		}
		result = max(result, v)
	}
	return result

}
