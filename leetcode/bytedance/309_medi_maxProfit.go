package main

//309. 最佳买卖股票时机含冷冻期
//给定一个整数数组，其中第i个元素代表了第i天的股票价格 。​
//
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//示例:
//
//输入: [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

//思路 单调递增 ，动态规划

//dp i 累计的当前最大钱数 状态：持有股票，不持有冻结中，不持有不冻结
// dp[0][持有] = -prices[0]，dp[0][冻结]=0，dp[0][不冻结]=0
// 当前持有时，前一个不可能是冻结 ,dp[i][持有] = max(dp[i-1][持有],d[i-1][不冻结]-prices[i])
// 不持有冻结中，dp[i][冻结] = dp[i-1][持有]+prices[i]
// 不持有不冻结，dp[i][不冻结] = max(dp[i-1][冻结],dp[i-1][不冻结])
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	dp := make([][3]int, n)

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(maxProfit([]int{8, 6, 4, 3, 3, 2, 3, 5, 8, 3, 8, 2, 6}))
}
