package main

import "sort"

//698. 划分为k个相等的子集
//给定一个整数数组nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
//
//
//
//示例 1：
//
//输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
//输出： True
//说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
//示例 2:
//
//输入: nums = [1,2,3,4], k = 3
//输出: false
//
//
//提示：
//
//1 <= k <= len(nums) <= 16
//0 < nums[i] < 10000
//每个元素的频率在 [1,4] 范围内

func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k > 0 {
		return false
	}
	per := sum / k
	sort.Ints(nums)
	if nums[n-1] > per {
		return false
	}
	dp := make([]bool, 1<<n)
	for i := range dp {
		dp[i] = true
	}
	var dfs func(int, int) bool
	dfs = func(s, p int) bool {
		if s == 0 {
			return true
		}
		if !dp[s] {
			return dp[s]
		}
		dp[s] = false
		for i, v := range nums {
			if v+p > per {
				break
			}
			if (s>>i)&1 > 0 && dfs(s^(1<<i), (v+p)%per) {
				return true
			}
		}
		return false
	}

	return dfs((1<<n)-1, 0)

}

func main() {
	println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}
