package main

//我们在两条独立的水平线上按给定的顺序写下A和B中的整数。
//
//现在，我们可以绘制一些连接两个数字A[i]和B[j]的直线，只要A[i] == B[j]，且我们绘制的直线不与任何其他连线（非水平线）相交。
//
//以这种方法绘制线条，并返回我们可以绘制的最大连线数。

//设 A[0] ~ A[x] 与 B[0] ~ B[y] 的最大连线数为 f(x, y)，那么对于任意位置的 f(i, j) 而言：
//如果 A[i] == B[j]，即 A[i] 和 B[j] 可连线，此时 f(i, j) = f(i - 1, j - 1) + 1
//如果 A[i] != B[j]，即 A[i] 和 B[j] 不可连线，此时最大连线数取决于 f(i - 1, j) 和 f(i, j - 1) 的较大值

func maxUncrossedLines(A []int, B []int) int {

	dp := [500][500]int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[len(A)][len(B)]

}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
