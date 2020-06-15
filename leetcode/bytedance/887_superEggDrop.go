package main

import "math"

// 复杂度 O(KN^2)
func superEggDropN2(K int, N int) int {
	type index struct {
		k int
		n int
	}
	memo := make(map[index]int)

	var dp func(k, n int) int
	dp = func(k, n int) int {
		if n == 0 {
			return 0
		}
		if k == 1 {
			return n
		}
		result := math.MaxInt32
		// 已经计算过的值
		if v, ok := memo[index{k, n}]; ok {
			return v
		}
		for i := 1; i <= n; i++ {
			result = min(result, 1+max(dp(k-1, i-1), dp(k, n-i)))
		}
		memo[index{k, n}] = result
		return result
	}
	return dp(K, N)

}

// 时间复杂度 O(KNlogN)

func superEggDrop(K int, N int) int {
	type index struct {
		k int
		n int
	}
	memo := make(map[index]int)

	var dp func(k, n int) int
	dp = func(k, n int) int {
		if n == 0 {
			return 0
		}
		if k == 1 {
			return n
		}
		result := math.MaxInt32
		if v, ok := memo[index{k, n}]; ok {
			return v
		}
		low, high := 1, n
		for low <= high {
			mid := (low + high) / 2
			broken := dp(k-1, mid-1)
			notBroken := dp(k, n-mid)
			if broken > notBroken {
				high = mid - 1
				result = min(result, broken+1)
			} else {
				low = mid + 1
				result = min(result, notBroken+1)
			}
		}

		memo[index{k, n}] = result
		return result
	}
	return dp(K, N)

}

func superEggDropBest(K, N int) int {
	moves := 0
	dp := make([]int, K+1) // 1 <= K <= 100
	// dp[i] = n 表示， i 个鸡蛋，利用 moves 次移动，最多可以检测 n 层楼
	for dp[K] < N {
		for i := K; i > 0; i-- {
			//逆序从K---1,dp[i] = dp[i]+dp[i-1] + 1 相当于上次移动后的结果,dp[]函数要理解成抽象出来的一个黑箱子函数,跟上一次移动时鸡蛋的结果有关系
			dp[i] += dp[i-1] + 1
			// 以上计算式，是从以下转移方程简化而来
			// dp[moves][k] = 1 + dp[moves-1][k-1] + dp[moves-1][k]
			// 假设 dp[moves-1][k-1] = n0, dp[moves-1][k] = n1
			// 首先检测，从第 n0+1 楼丢下鸡蛋会不会破。
			// 如果鸡蛋破了，F 一定是在 [1：n0] 楼中，
			//         利用剩下的 moves-1 次机会和 k-1 个鸡蛋，可以把 F 找出来。
			// 如果鸡蛋没破，假如 F 在 [n0+2:n0+n1+1] 楼中
			//         利用剩下的 moves-1 次机会和 k 个鸡蛋把，也可以把 F 找出来。
			// 所以，当有 moves 个放置机会和 k 个鸡蛋的时候
			// F 在 [1, n0+n1+1] 中的任何一楼，都能够被检测出来。
		}
		moves++
	}
	return moves
}

func superEggDropBest2(K, N int) int {
	moves := 0
	dp := make([]int, K+1) // 1 <= K <= 100
	for dp[K] < N {
		for i := K; i > 0; i-- {
			dp[i] += dp[i-1] + 1
		}
		moves++
	}
	return moves
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(superEggDrop(2, 6))
}
