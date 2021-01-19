package main

import "sort"

//1584. 连接所有点的最小费用
//给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。
//
//连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示 val 的绝对值。
//
//请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。
//
//示例 2：
//
//输入：points = [[3,12],[-2,5],[-4,1]]
//输出：18
//示例 3：
//
//输入：points = [[0,0],[1,1],[1,0],[-1,1]]
//输出：4
//示例 4：
//
//输入：points = [[-1000000,-1000000],[1000000,1000000]]
//输出：4000000
//示例 5：
//
//输入：points = [[0,0]]
//输出：0
// 
//
//提示：
//
//1 <= points.length <= 1000
//-106 <= xi, yi <= 106
//所有点 (xi, yi) 两两不同。

//思路 Kruskal算法 ,并查集
// Kruskal算法是最小生成树的经典算法
// Kruskal算法是从小到大加入边的贪心算法
func minCostConnectPoints(points [][]int) int {
	edges := make([]edge, 0)
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{
				x:    i,
				y:    j,
				cost: abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1]),
			})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})
	parent := make([]int, n)
	count := make([]int, n)
	for i := range parent {
		parent[i] = i
		count[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) bool {
		from, to = find(from), find(to)
		if from == to{
			return false
		}
		if count[from] > count[to]{
			from,to = to,from
		}
		count[to] += count[from]
		parent[from] = to
		return true
	}

	var result int
	for _, v := range edges {
		if union(v.x,v.y){
			result += v.cost
		}
	}
	return result
}

type edge struct {
	x, y int
	cost int
}
