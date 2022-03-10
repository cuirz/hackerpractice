package main

import (
	"math"
)

//322. 零钱兑换
//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回-1。

//示例1:
//
//输入: coins = [1, 2, 5], amount = 11
//输出: 3
//解释: 11 = 5 + 5 + 1
//示例 2:
//
//输入: coins = [2], amount = 3
//输出: -1

//方案 动态规划 自上而下
// dp(S) = dp(S-C)+1

func coinChange(coins []int, amount int) int {
	count := make([]int, amount)
	var dp func(s int) int
	dp = func(s int) int {
		if s < 0 {
			return -1
		}
		if s == 0 {
			return 0
		}
		if count[s-1] > 0 {
			return count[s-1]
		}

		mini := math.MaxInt32
		for _, v := range coins {
			temp := dp(s - v)
			if temp >= 0 && temp < mini {
				mini = temp + 1
			}
		}
		if mini == math.MaxInt32 {
			return -1
		}
		count[s-1] = mini
		return mini
	}

	return dp(amount)

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}

func main() {
	println(coinChange([]int{1, 2, 5}, 11))
}
