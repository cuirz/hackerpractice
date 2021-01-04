package main

import "math"

//188. 买卖股票的最佳时机 IV
//给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 
//
//示例 1：
//
//输入：k = 2, prices = [2,4,1]
//输出：2
//解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
//示例 2：
//
//输入：k = 2, prices = [3,2,6,5,0,3]
//输出：7
//解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
// 
//
//提示：
//
//0 <= k <= 109
//0 <= prices.length <= 1000
//0 <= prices[i] <= 1000

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := make([][]int, n)
	sell := make([][]int, n)
	k = min(k, n/2)
	for i, _ := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i], sell[0][i] = math.MinInt32>>1, math.MinInt32>>1
	}
	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])

		}
	}
	return max(sell[n-1]...)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x ...int) int {
	res := x[0]
	for _, v := range x[1:] {
		if res < v {
			res = v
		}
	}
	return res
}
