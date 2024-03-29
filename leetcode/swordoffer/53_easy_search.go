package main

import "sort"

//剑指 Offer 53 - I. 在排序数组中查找数字 I
//统计一个数字在排序数组中出现的次数。
//
//
//
//示例 1:
//
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: 2
//示例2:
//
//输入: nums = [5,7,7,8,8,10], target = 6
//输出: 0
//
//
//限制：
//
//0 <= 数组长度 <= 50000

func search(nums []int, target int) int {
	index := sort.SearchInts(nums, target)
	n := len(nums)
	result := 0
	for i := index; i < n; i++ {
		if nums[i] == target {
			result++
		}
	}
	for i := index - 1; i > -1; i-- {
		if nums[i] == target {
			result++
		}
	}
	return result
}
func main() {
	println(search([]int{5, 7, 7, 8, 8, 10}, 6))
}
