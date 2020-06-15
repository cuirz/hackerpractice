package main

//862. 和至少为 K 的最短子数组

// 前缀和 ，双端队列 ，单调递增

func shortestSubarray(A []int, K int) int {
	n := len(A)

	result := n + 1
	sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + A[i]
	}

	queue := make([]int, 0)
	for i := 0; i < len(sums); i++ {
		for len(queue) > 0 && sums[i] <= sums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		for len(queue) > 0 && sums[i]-sums[queue[0]] >= K {
			result = min(result, i-queue[0])
			queue = queue[1:]
		}
		queue = append(queue, i)
	}
	if result == n+1 {
		result = -1
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	//println(shortestSubarray([]int{84, -37, 32, 40, 95}, 167))
	println(shortestSubarray([]int{2, -1, 2}, 3))
	//[84,-37,32,40,95]
	//167
}
