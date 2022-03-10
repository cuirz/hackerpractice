package main

import "sort"

//368. 最大整除子集
//给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[j]) 都应当满足：
//answer[i] % answer[j] == 0 ，或
//answer[j] % answer[i] == 0
//如果存在多个有效解子集，返回其中任何一个均可。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[1,2]
//解释：[1,3] 也会被视为正确答案。
//示例 2：
//
//输入：nums = [1,2,4,8]
//输出：[1,2,4,8]
//
//
//提示：
//
//1 <= nums.length <= 1000
//1 <= nums[i] <= 2 * 109
//nums 中的所有整数 互不相同

//思路 动态规划
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	maxCount, maxVal := 1, 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxCount {
			maxCount, maxVal = dp[i], nums[i]
		}
	}
	if maxCount == 1 {
		return []int{nums[0]}
	}
	result := make([]int, 0)
	for i := n - 1; i > -1 && maxCount > 0; i-- {
		if dp[i] == maxCount && maxVal%nums[i] == 0 {
			result = append(result, nums[i])
			maxVal = nums[i]
			maxCount--
		}
	}
	return result
}
