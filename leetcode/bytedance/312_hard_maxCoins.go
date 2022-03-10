package main

//312. 戳气球
//有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组nums中。
//
//现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得nums[left] * nums[i] * nums[right]个硬币。这里的left和right代表和i相邻的两个气球的序号。注意当你戳破了气球 i 后，气球left和气球right就变成了相邻的气球。
//
//求所能获得硬币的最大数量。
//
//说明:
//
//你可以假设nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
//0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
//示例:
//
//输入: [3,1,5,8]
//输出: 167
//解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//    coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

//思路 动态规划，分治法
//倒序 区间[i,j] 里 最后一个戳破的气球假设为 k，那最后值为 v[i]*v[j]*v[k]
//戳破k之前就分为两个区间  [i,k]-[k,j], 每个区间循环上一操作
//dp[i][j] = max(v[i]*v[k][j] + dp[i][k]+dp[k][j])

func maxCoins(nums []int) int {
	n := len(nums) + 2
	dp := make([][]int, n)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n-1; i++ {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[j]*nums[k])
			}
		}
	}
	return dp[0][n-1]
}
