package main

import "sort"

//1365. 有多少小于当前数字的数字
//给你一个数组nums，对于其中每个元素nums[i]，请你统计数组中比它小的所有数字的数目。
//
//换而言之，对于每个nums[i]你必须计算出有效的j的数量，其中 j 满足j != i 且 nums[j] < nums[i]。

func smallerNumbersThanCurrent(nums []int) []int {
	dic := make(map[int][]int)
	for i, v := range nums {
		dic[v] = append(dic[v], i)
	}
	sort.Ints(nums)
	result := make([]int, len(nums))
	var pre, index int
	for i, v := range nums {
		if pre != v {
			pre = v
			index = i
			for _, val := range dic[v] {
				result[val] = index
			}
		}
	}
	return result
}
