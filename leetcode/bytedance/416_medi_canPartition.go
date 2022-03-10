package main

//416. 分割等和子集
//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//注意:
//
//每个数组中的元素不会超过 100
//数组的大小不会超过 200
//示例 1:
//
//输入: [1, 5, 11, 5]
//
//输出: true
//
//解释: 数组可以分割成 [1, 5, 5] 和 [11].
//
//
//示例2:
//
//输入: [1, 2, 3, 5]
//
//输出: false
//
//解释: 数组不能分割成两个元素和相等的子集.

//思路  动态规划
// 背包问题
// 判断能不能装满背包  = 判断数组中子集和等于target
// dp[i][j] ,i nums数组下标，j target下标
// 装完第i个，j容量 状态
// 已经装满时，dp[i][j] = dp[i-1][j]
// 没有装满时,放入或不放入 状态转移，dp[i][j] = dp[i-1][j] || dp[i][j-nums[i-1]]

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum >> 1
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][target]
}
