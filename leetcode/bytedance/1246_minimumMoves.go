package main

import "math"

//f(i,j)
// if arr[i] == arr[j]   f(i+1,j-1)
// min(f(i,j), f(i,k)+f(k+1,j))
func minimumMoves(arr []int) int {
	size := len(arr)
	if size < 1 {
		return 0
	}
	var dp [100][100]int

	for i := 0; i < size; i++ {
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if j == i+1 {
				if arr[i] == arr[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = 2
				}
				continue
			}
			temp := math.MaxInt32
			for k := i; k < j; k++ {
				temp = min(temp, dp[i][k]+dp[k+1][j])
			}

			if arr[i] == arr[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = temp
			}
		}
	}
	return dp[0][size-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(minimumMoves([]int{1, 3, 4, 1, 5}))
}
