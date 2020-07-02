package main

//718. 最长重复子数组
//给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
//
//示例 1:
//
//输入:
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//输出: 3
//解释:
//长度最长的公共子数组是 [3, 2, 1]。
//说明:
//
//1 <= len(A), len(B) <= 1000
//0 <= A[i], B[i] < 100

// 思路： 动态规划, 滑动窗口
// 当已知最长公共子数组是 A[i:] = B[j:]时， dp[i][j] = dp[i+1][j+1] + 1
// 为了先计算得到 dp[i+1][j+1]值，倒序数组进行比较
func findLength(A []int, B []int) int {
	an := len(A)
	bn := len(B)
	dp := make([][]int, an+1)
	result := 0
	for i := 0; i < an+1; i++ {
		dp[i] = make([]int, bn+1)
	}
	for i := an - 1; i > -1; i-- {
		for j := bn - 1; j > -1; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				result = max(result, dp[i][j])
			}
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(findLength([]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}))
}
