package main

//862. 和至少为 K 的最短子数组
//给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回 -1 。
//
//子数组 是数组中 连续 的一部分。
//
//
//
//示例 1：
//
//输入：nums = [1], k = 1
//输出：1
//示例 2：
//
//输入：nums = [1,2], k = 4
//输出：-1
//示例 3：
//
//输入：nums = [2,-1,2], k = 3
//输出：3
//
//
//提示：
//
//1 <= nums.length <= 10^5
//-10^5 <= nums[i] <= 10^5
//1 <= k <= 10^9

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}

	result := n + 1
	var q []int
	for i, v := range preSum {
		for len(q) > 0 && v-preSum[q[0]] >= k {
			result = min(result, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && preSum[q[len(q)-1]] >= v {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if result < n+1 {
		return result
	}
	return -1

}
