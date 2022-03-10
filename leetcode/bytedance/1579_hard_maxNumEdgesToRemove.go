package main

import "sort"

//1579. 保证图可完全遍历
//Alice 和 Bob 共有一个无向图，其中包含 n 个节点和 3 种类型的边：
//
//类型 1：只能由 Alice 遍历。
//类型 2：只能由 Bob 遍历。
//类型 3：Alice 和 Bob 都可以遍历。
//给你一个数组 edges ，其中 edges[i] = [typei, ui, vi]表示节点 ui 和 vi 之间存在类型为 typei 的双向边。请你在保证图仍能够被 Alice和 Bob 完全遍历的前提下，找出可以删除的最大边数。如果从任何节点开始，Alice 和 Bob 都可以到达所有其他节点，则认为图是可以完全遍历的。
//
//返回可以删除的最大边数，如果 Alice 和 Bob 无法完全遍历图，则返回 -1 。
//
//输入：n = 4, edges = [[3,2,3],[1,1,2],[2,3,4]]
//输出：-1
//解释：在当前图中，Alice 无法从其他节点到达节点 4 。类似地，Bob 也不能达到节点 1 。因此，图无法完全遍历。
//
//
//提示：
//
//1 <= n <= 10^5
//1 <= edges.length <= min(10^5, 3 * n * (n-1) / 2)
//edges[i].length == 3
//1 <= edges[i][0] <= 3
//1 <= edges[i][1] < edges[i][2] <= n
//所有元组 (typei, ui, vi) 互不相同

//思路 并查集
func maxNumEdgesToRemove(n int, edges [][]int) int {

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] > edges[j][0]
	})

	parent := make([][2]int, n+1)
	size := make([][2]int, n+1)

	for i := range parent {
		parent[i][0] = i
		parent[i][1] = i
		size[i][0] = 1
		size[i][1] = 1
	}

	var find func(x, t int) int
	find = func(x, t int) int {
		if parent[x][t] != x {
			parent[x][t] = find(parent[x][t], t)
		}
		return parent[x][t]
	}

	union := func(x, y, t int) bool {
		x, y = find(x, t), find(y, t)
		if x == y {
			return false
		}
		if size[x][t] < size[y][t] {
			x, y = y, x
		}
		size[x][t] += size[y][t]
		parent[y][t] = x
		return true
	}

	result := 0
	for _, v := range edges {
		t1, t2 := false, false
		if v[0] == 1 || v[0] == 3 {
			t1 = union(v[1], v[2], 0)
		}
		if v[0] == 2 || v[0] == 3 {
			t2 = union(v[1], v[2], 1)
		}
		if !t1 && !t2 {
			result++
		}
	}

	for t := 0; t < 2; t++ {
		p := find(1, t)
		for i := 2; i <= n; i++ {
			if p != find(i, t) {
				return -1
			}
		}
	}
	return result
}
