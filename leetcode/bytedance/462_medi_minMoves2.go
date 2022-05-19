package main

import "sort"

//462. 最少移动次数使数组元素相等 II
//给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最少移动数。
//
//在一步操作中，你可以使数组中的一个元素加 1 或者减 1 。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：2
//解释：
//只需要两步操作（每步操作指南使一个元素加 1 或减 1）：
//[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
//示例 2：
//
//输入：nums = [1,10,2,9]
//输出：16
//
//提示：
//
//n == nums.length
//1 <= nums.length <= 10^5
//-10^9 <= nums[i] <= 10^9

func minMoves2(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	mid := nums[n/2]
	var result int
	for _, v := range nums {
		result += abs(v - mid)
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
