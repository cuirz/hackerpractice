package main

func longestSubsequence(arr []int, difference int) int {
	n := len(arr)
	dp := make(map[int]int, n)
	result := 1
	for _, v := range arr {
		pre := v - difference
		if dp[pre] > 0 {
			dp[v] = dp[pre] + 1
			result = max(result, dp[v])
		} else {
			dp[v] = 1
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
	println(longestSubsequence([]int{1, 2, 3, 4}, 1))
}
