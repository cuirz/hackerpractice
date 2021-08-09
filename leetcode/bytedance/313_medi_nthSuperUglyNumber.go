package main

import "math"

//313. 超级丑数
//超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。
//
//给你一个整数 n 和一个整数数组 primes ，返回第 n 个 超级丑数 。
//
//题目数据保证第 n 个 超级丑数 在 32-bit 带符号整数范围内。
//
// 
//
//示例 1：
//
//输入：n = 12, primes = [2,7,13,19]
//输出：32
//解释：给定长度为 4 的质数数组 primes = [2,7,13,19]，前 12 个超级丑数序列为：[1,2,4,7,8,13,14,16,19,26,28,32] 。
//示例 2：
//
//输入：n = 1, primes = [2,3,5]
//输出：1
//解释：1 不含质因数，因此它的所有质因数都在质数数组 primes = [2,3,5] 中。
// 
//提示：
//
//1 <= n <= 106
//1 <= primes.length <= 100
//2 <= primes[i] <= 1000
//题目数据 保证 primes[i] 是一个质数
//primes 中的所有值都 互不相同 ，且按 递增顺序 排列

//思路  最小堆，第n次取出来的就是 第n个超级丑数
// 动态规划
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	m := len(primes)
	points := make([]int, m)
	for i := range primes {
		points[i] = 1
	}
	for i := 2; i <= n; i++ {
		minNum := math.MaxInt32
		nums := make([]int, m)
		for j := 0; j < m; j++ {
			nums[j] = dp[points[j]] * primes[j]
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j, v := range nums {
			if v == minNum {
				points[j] ++
			}
		}
	}
	return dp[n]
}
