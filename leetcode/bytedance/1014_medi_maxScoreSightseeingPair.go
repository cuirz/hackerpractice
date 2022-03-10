package main

//1014. 最佳观光组合
//给定正整数数组A，A[i]表示第 i 个观光景点的评分，并且两个景点i 和j之间的距离为j - i。
//
//一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i- j）：景点的评分之和减去它们两者之间的距离。
//
//返回一对观光景点能取得的最高分。
//
//示例：
//
//输入：[8,1,5,2,6]
//输出：11
//解释：i = 0, j = 2, A[i] + A[j] + i - j = 8 + 5 + 0 - 2 = 11

//思路： 动态规划,枚举
// dp[j] = max(dp[i],max(A[i...j])+A[j]-j)
func maxScoreSightseeingPair(A []int) int {
	n := len(A)
	dp := make([]int, n)
	dp[0] = A[0]
	maxv := A[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], maxv+A[i]-i)
		maxv = max(maxv, A[i]+i)
	}
	return dp[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
}
