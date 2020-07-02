package main

import "sort"

//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
//注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//示例 1:
//
//输入: [2,4,1], k = 2
//输出: 2
//解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
//示例 2:
//
//输入: [3,2,6,5,0,3], k = 2
//输出: 7
//解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。

//思路 单调
func maxProfit(k int, prices []int) int {
	result := make([]int, 0)
	n := len(prices)
	if n < 1 {
		return 0
	}
	maxv := 0
	pre := prices[0]
	for i := 1; i < n; i++ {
		tmp := prices[i] - pre
		if maxv < tmp {
			maxv = tmp
			if i == n-1 {
				result = append(result, maxv)
			}
		} else {
			pre = prices[i]
			if maxv > 0 {
				result = append(result, maxv)
				maxv = 0
			}
		}
	}
	sort.Ints(result)
	res := 0
	n = max(len(result)-k-1, -1)
	for i := len(result) - 1; i > n; i-- {
		res += result[i]
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(maxProfit(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}
