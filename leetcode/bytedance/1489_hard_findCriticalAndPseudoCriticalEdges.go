package main

import "sort"

//1489. 找到最小生成树里的关键边和伪关键边
//给你一个 n个点的带权无向连通图，节点编号为 0到 n-1，同时还有一个数组 edges，其中 edges[i] = [fromi, toi, weighti]表示在fromi和toi节点之间有一条带权无向边。最小生成树(MST) 是给定图中边的一个子集，它连接了所有节点且没有环，而且这些边的权值和最小。
//
//请你找到给定图中最小生成树的所有关键边和伪关键边。如果从图中删去某条边，会导致最小生成树的权值和增加，那么我们就说它是一条关键边。伪关键边则是可能会出现在某些最小生成树中但不会出现在所有最小生成树中的边。
//
//请注意，你可以分别以任意顺序返回关键边的下标和伪关键边的下标。
//
//
//
//示例 1：
//
//
//
//输入：n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
//输出：[[0,1],[2,3,4,5]]
//解释：上图描述了给定图。
//下图是所有的最小生成树。
//
//注意到第 0 条边和第 1 条边出现在了所有最小生成树中，所以它们是关键边，我们将这两个下标作为输出的第一个列表。
//边 2，3，4 和 5 是所有 MST 的剩余边，所以它们是伪关键边。我们将它们作为输出的第二个列表。
//示例 2 ：
//
//
//
//输入：n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]
//输出：[[],[0,1,2,3]]
//解释：可以观察到 4 条边都有相同的权值，任选它们中的 3 条可以形成一棵 MST 。所以 4 条边都是伪关键边。
//
//
//提示：
//
//2 <= n <= 100
//1 <= edges.length <= min(200, n * (n - 1) / 2)
//edges[i].length == 3
//0 <= fromi < toi < n
//1 <= weighti<= 1000
//所有 (fromi, toi)数对都是互不相同的。

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	type connect struct {
		next int
		eid  int
	}
	size := len(edges)
	for i := range edges {
		edges[i] = append(edges[i], i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	edgeType := make([]int, size)
	graph := make([][]connect, n)
	dfn := make([]int, n)
	timestamp := 0
	var tarjan func(v, id int) int
	tarjan = func(v, id int) int {
		timestamp++
		dfn[v] = timestamp
		lowV := timestamp
		for _, con := range graph[v] {
			if w := con.next; dfn[w] == 0 {
				lowW := tarjan(w, con.eid)
				if dfn[v] < lowW {
					edgeType[con.eid] = 1
				}
				lowV = min(lowV, lowW)
			} else if con.eid != id {
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}

	us := newUnionSet(n)
	for i := 0; i < size; {
		nodes := make([]int, 0)
		for weight := edges[i][2]; i < size && weight == edges[i][2]; i++ {
			v, w := us.find(edges[i][0]), us.find(edges[i][1])
			eid := edges[i][3]
			if v != w {
				graph[v] = append(graph[v], connect{w, eid})
				graph[w] = append(graph[w], connect{v, eid})
				nodes = append(nodes, v, w)
			} else {
				edgeType[eid] = -1
			}
		}
		for _, v := range nodes {
			if dfn[v] == 0 {
				tarjan(v, -1)
			}
		}
		for i := 0; i < len(nodes); i += 2 {
			v, w := nodes[i], nodes[i+1]
			us.union(v, w)
			graph[v], graph[w] = nil, nil
			dfn[v], dfn[w] = 0, 0
		}
	}
	result := make([][]int, 2)
	for i, v := range edgeType {
		switch v {
		case 1:
			result[0] = append(result[0], i)
		case 0:
			result[1] = append(result[1], i)
		}
	}
	return result
}

type UnionSet struct {
	Parent []int
	Count  []int
}

func newUnionSet(n int) *UnionSet {
	p, c := make([]int, n), make([]int, n)
	for i := range p {
		p[i] = i
		c[i] = 1
	}
	return &UnionSet{p, c}
}

func (u *UnionSet) find(x int) int {
	if u.Parent[x] != x {
		u.Parent[x] = u.find(u.Parent[x])
	}
	return u.Parent[x]

}
func (u *UnionSet) union(x, y int) bool {
	x, y = u.find(x), u.find(y)
	if x == y {
		return false
	}
	if u.Count[x] < u.Count[y] {
		x, y = y, x
	}
	u.Count[x] += u.Count[y]
	u.Parent[y] = x
	return true
}

func main() {
	//findCriticalAndPseudoCriticalEdges(5, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}})
	findCriticalAndPseudoCriticalEdges(6, [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 1}, {2, 3, 4}, {3, 4, 2}, {3, 5, 2}, {4, 5, 2}})
}
