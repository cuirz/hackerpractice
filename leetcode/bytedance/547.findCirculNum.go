package main

func findCircleNum(M [][]int) int {
	dp := make(map[int]bool, 0)
	result := 0
	for i := 0; i < len(M); i++ {
		if dp[i] {
			continue
		}
		dp[i] = true
		result++
		check(M, dp, i)
	}
	return result
}

func check(M [][]int, dp map[int]bool, index int) {
	for i := 0; i < len(M); i++ {
		if i != index && M[index][i] == 1 && !dp[i] {
			dp[i] = true
			check(M, dp, i)
		}
	}

}
