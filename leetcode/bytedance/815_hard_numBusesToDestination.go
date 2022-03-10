package main

import "math"

//815. 公交路线
//给你一个数组 routes ，表示一系列公交线路，其中每个 routes[i] 表示一条公交线路，第 i 辆公交车将会在上面循环行驶。
//
//例如，路线 routes[0] = [1, 5, 7] 表示第 0 辆公交车会一直按序列 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... 这样的车站路线行驶。
//现在从 source 车站出发（初始时不在公交车上），要前往 target 车站。 期间仅可乘坐公交车。
//
//求出 最少乘坐的公交车数量 。如果不可能到达终点车站，返回 -1 。
//
//
//
//示例 1：
//
//输入：routes = [[1,2,7],[3,6,7]], source = 1, target = 6
//输出：2
//解释：最优策略是先乘坐第一辆公交车到达车站 7 , 然后换乘第二辆公交车到车站 6 。
//示例 2：
//
//输入：routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
//输出：-1
//
//
//提示：
//
//1 <= routes.length <= 500.
//1 <= routes[i].length <= 10^5
//routes[i] 中的所有值 互不相同
//sum(routes[i].length) <= 10^5
//0 <= routes[i][j] < 10^6
//0 <= source, target < 10^6

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	nodes := make(map[int][]int)
	n := len(routes)
	edge := make([][]bool, n)
	for i := range edge {
		edge[i] = make([]bool, n)
	}

	for i, r := range routes {
		for _, v := range r {
			for _, j := range nodes[v] {
				edge[i][j], edge[j][i] = true, true
			}
			nodes[v] = append(nodes[v], i)
		}
	}
	dis := make([]int, n)
	queue := make([]int, 0)
	for _, v := range nodes[source] {
		dis[v] = 1
		queue = append(queue, v)
	}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for j := 0; j < n; j++ {
			if edge[i][j] && dis[j] == 0 {
				dis[j] = dis[i] + 1
				queue = append(queue, j)
			}
		}
	}
	result := math.MaxInt32
	for _, v := range nodes[target] {
		if dis[v] > 0 {
			result = min(result, dis[v])
		}
	}
	if result == math.MaxInt32 {
		result = -1
	}
	return result
}

func main() {
	println(numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12))
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
