package main

//17.21. 直方图的水量
//给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为
//输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6

//思路 动态规划
func trap(height []int) int {
	n := len(height)
	dp := make([]int, n)
	var spike int
	for i := 0; i < n; i++ {
		spike = max(spike, height[i])
		dp[i] = spike
	}
	var result int
	spike = 0
	for i := n - 1; i > -1; i-- {
		spike = max(spike, height[i])
		val := min(spike, dp[i]) - height[i]
		if val > 0 {
			result += val
		}
	}
	return result
}
