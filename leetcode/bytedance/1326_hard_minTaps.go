package main

import "math"

//
//在 x 轴上有一个一维的花园。花园长度为 n，从点 0 开始，到点 n 结束。
//
//花园里总共有 n + 1 个水龙头，分别位于 [0, 1, ..., n] 。
//
//给你一个整数 n 和一个长度为 n + 1 的整数数组 ranges ，其中 ranges[i] （下标从 0 开始）表示：如果打开点 i 处的水龙头，可以灌溉的区域为 [i -  ranges[i], i + ranges[i]] 。
//
//请你返回可以灌溉整个花园的 最少水龙头数目 。如果花园始终存在无法灌溉到的地方，请你返回 -1 。
//
//
//
//示例 1：
//输入：n = 5, ranges = [3,4,1,1,0,0]
//输出：1
//解释：
//点 0 处的水龙头可以灌溉区间 [-3,3]
//点 1 处的水龙头可以灌溉区间 [-3,5]
//点 2 处的水龙头可以灌溉区间 [1,3]
//点 3 处的水龙头可以灌溉区间 [2,4]
//点 4 处的水龙头可以灌溉区间 [4,4]
//点 5 处的水龙头可以灌溉区间 [5,5]
//只需要打开点 1 处的水龙头即可灌溉整个花园 [0,5] 。

// 方案  贪婪算法
// 指定一个【0，n] result区间，上面记录区间范围
// 从水龙头【l,r] r值相同时 记录l 最大
// 从 n 尾部开始递减扫result，

func minTaps(n int, ranges []int) int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		l := max(i-ranges[i], 0)
		r := min(i+ranges[i], n)
		result[l] = max(result[l], r)
	}

	count := 0
	best := math.MinInt32
	seek := 0
	for i := 0; i < n; i++ {
		best = max(best, result[i])
		if i == seek {
			if best <= i {
				return -1
			}
			seek = best
			count++
		}
	}
	return count
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
