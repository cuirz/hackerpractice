package main

//713. 乘积小于K的子数组
//给定一个正整数数组 nums。
//
//找出该数组内乘积小于 k 的连续的子数组的个数。
//输入: nums = [10,5,2,6], k = 100
//输出: 8
//解释: 8个乘积小于100的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
//需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
//
//说明:
//
//0 < nums.length <= 50000
//0 < nums[i] < 1000
//0 <= k < 10^6

// 方法 ： 双指针

func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	left := 0
	sums := 1
	for right, v := range nums {
		for sums *= v; sums > k; {
			sums /= nums[left]
			left++
		}
		count += right - left + 1
	}
	return count

}
