package main

//787. K 站中转内最便宜的航班
//有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城市 fromi 开始，以价格 toi 抵达 pricei。
//
//现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
//
// 
//
//示例 1：
//
//输入:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 1
//输出: 200
//解释:
//城市航班图如下
//
//
//从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
//示例 2：
//
//输入:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 0
//输出: 500
//解释:
//城市航班图如下
//
//
//从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。
// 
//
//提示：
//
//1 <= n <= 100
//0 <= flights.length <= (n * (n - 1) / 2)
//flights[i].length == 3
//0 <= fromi, toi < n
//fromi != toi
//1 <= pricei <= 10^4
//航班没有重复，且不存在自环
//0 <= src, dst, k < n
//src != dst

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 1000*101+1
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][src] = 0
	for c := 1; c <= k+1; c++ {
		for _, f := range flights {
			s, d, cost := f[0], f[1], f[2]
			dp[c][d] = min(dp[c][d], dp[c-1][s]+cost)
		}
	}
	result := inf
	for c := 1; c <= k+1; c++ {
		result = min(result, dp[c][dst])
	}
	if result == inf{
		return -1
	}
	return result

}
