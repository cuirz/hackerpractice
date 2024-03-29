package main

import "sort"

//347. 前 K 个高频元素
//给定一个非空的整数数组，返回其中出现频率前k高的元素。
//
//
//
//示例 1:
//
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//示例 2:
//
//输入: nums = [1], k = 1
//输出: [1]
//
//
//提示：
//
//你可以假设给定的k总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
//你的算法的时间复杂度必须优于 O(n log n) ,n是数组的大小。
//题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
//你可以按任意顺序返回答案。

func topKFrequent(nums []int, k int) []int {
	dic := make(map[int]int)
	array := make([]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if dic[nums[i]] == 0 {
			array = append(array, nums[i])
		}
		dic[nums[i]]++
	}
	sort.Slice(array, func(i, j int) bool {
		return dic[array[i]] > dic[array[j]]
	})
	return array[:k]
}

func main() {
	println(topKFrequent([]int{1, 1, 1, 1, 2, 2, 2, 2, 3}, 2))
}
