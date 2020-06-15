package main

//if not matrix:
//return 0
//rows, cols, max_size = len(matrix), len(matrix[0]), 0
//size = [0] * cols
//for i in range(rows):
//count = 0
//for j in range(cols):
//size[j] += 1 if matrix[i][j] == '1' else 0
//for j in range(cols):
//if size[j] > max_size:
//count += 1
//if count > max_size:
//max_size += 1
//break
//else:
//count = 0
//return max_size * max_size
func maximalSquare(matrix [][]byte) int {
	if len(matrix) < 1 {
		return 0
	}
	rows, cols, max_size := len(matrix), len(matrix[0]), 0
	dp := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				}
				if dp[i][j] > max_size {
					max_size = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return max_size * max_size
}

func min(x, y, z int) int {
	if x > y {
		if y > z {
			return z
		}
		return y
	}
	if x > z {
		return z
	}
	return x
}

func main() {
	//maximalSquare([["1","0","1","1","0","1"],["1","1","1","1","1","1"],["0","1","1","0","1","1"],["1","1","1","0","1","0"],["0","1","1","1","1","1"],["1","1","0","1","1","1"]])
}
