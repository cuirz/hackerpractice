package main

//17.10. 主要元素
//数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。
//
//
//
//示例 1：
//
//输入：[1,2,5,9,5,9,5,5,5]
//输出：5
//示例 2：
//
//输入：[3,2]
//输出：-1
//示例 3：
//
//输入：[2,2,1,1,1,2,2]
//输出：2

//思路 摩尔投票
func majorityElement(nums []int) int {
	candidate, count := -1, 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if candidate == v {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, v := range nums {
		if candidate == v {
			count++
		}
	}
	if count > len(nums)/2 {
		return candidate
	}
	return -1
}
