package main

//494. 目标和
//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号+和-。对于数组中的任意一个整数，你都可以从+或-中选择一个符号添加在前面。
//
//返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
//
//
//
//示例：
//
//输入：nums: [1, 1, 1, 1, 1], S: 3
//输出：5
//解释：
//
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
//
//一共有5种方法让最终目标和为3。
//
//
//提示：
//
//数组非空，且长度不会超过 20 。
//初始的数组的和不会超过 1000 。
//保证返回的最终结果能被 32 位整数存下。

//思路 枚举

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	result := 0
	var dfs func(sum, index int)
	dfs = func(sum, index int) {
		if index == n {
			if sum == S {
				result++
			}
			return
		}
		dfs(sum+nums[index], index+1)
		dfs(sum-nums[index], index+1)
	}
	dfs(0, 0)
	return result
}

//思路 动态规划
// 背包和的问题
// dp[i][S]
// 举例 [1,2,3,4,5]
// 把数 划分出 正子集 和 负子集
// 正[1,3,4]  负[2,5]
// sum(P) - sum(N) = target
// sum(nums) = sum(P) + sum(N)
// sum(P) = (target + sum(nums))/2
// 故 target + sum 肯定是偶数
func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum < abs(S) || (sum+S)%2 != 0 {
		return 0
	}

	positive := (sum + S) >> 1
	dp := make([]int, positive+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := positive; j >= nums[i]; j-- {
			dp[j] = dp[j] + dp[j-nums[i]]
		}
	}
	return dp[positive]
}

func abs(x int) int {
	if x < 0 {
		return x
	}
	return x

}
