package main

//1319. 连通网络的操作次数
//用以太网线缆将n台计算机连接成一个网络，计算机的编号从0到n-1。线缆用connections表示，其中connections[i} = [a, b]连接了计算机a和b。
//
//网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。
//
//给你这个计算机网络的初始布线connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回-1 。
//示例 3：
//
//输入：n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
//输出：-1
//解释：线缆数量不足。
//示例 4：
//
//输入：n = 5, connections = [[0,1],[0,2],[3,4],[2,3]]
//输出：0
//
//
//提示：
//
//1 <= n <= 10^5
//1 <= connections.length <= min(n*(n-1)/2, 10^5)
//connections[i].length == 2
//0 <= connections[i][0], connections[i][1]< n
//connections[i][0] != connections[i][1]
//没有重复的连接。
//两台计算机不会通过多条线缆连接。

//思路 并查集
// 寻找有多少独立的集，计算出多余的线
// 有几个独立集就需要相同数量的线
func makeConnected(n int, connections [][]int) int {

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
	union := func(x, y int) bool {
		x, y = find(x), find(y)
		if x == y {
			return false
		}
		if count[x] < count[y] {
			x, y = y, x
		}
		count[x] += count[y]
		parent[y] = x
		return true
	}

	surplus := 0
	for _, v := range connections {
		if !union(v[0], v[1]) {
			surplus++
		}
	}

	result := -1
	class := make(map[int]bool)
	for _, v := range parent {
		p := find(v)
		if !class[p] {
			class[p] = true
			result++
		}
	}
	if result > surplus {
		return -1
	}
	return result
}

func main() {
	//makeConnected(5,[][]int{{0,1},{0,2},{3,4},{2,3}})
	makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}})

}
