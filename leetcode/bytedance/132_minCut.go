package main

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = -1
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 0; j < i; j++ {
			if isPal(s, j, i-1) {
				dp[i] = Min(dp[j]+1, dp[i])
			}
		}
	}
	return dp[n]
}

func isPal(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(minCut("abac"))
}
