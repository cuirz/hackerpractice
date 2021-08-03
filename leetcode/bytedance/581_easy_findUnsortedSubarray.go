package main

import "math"

//581. 最短无序连续子数组
//给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
//你找到的子数组应是最短的，请输出它的长度。
//
//示例 1:
//
//输入: [2, 6, 4, 8, 10, 9, 15]
//输出: 5
//解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//说明 :
//
//输入的数组长度范围在 [1, 10,000]。
//输入的数组可能包含重复元素 ，所以升序的意思是<=。
func findUnsortedSubarray(nums []int) int {
	minv, maxv := math.MaxInt32, math.MinInt32
	l, r := -1, -1
	n := len(nums)
	for i, v := range nums {
		if maxv > v {
			r = i
		} else {
			maxv = v
		}
		if minv < nums[n-i-1] {
			l = n - i - 1
		} else {
			minv = nums[n-i-1]
		}

	}
	if r == -1 {
		return 0
	}
	return r - l + 1
}
func findUnsortedSubarray2(nums []int) int {

	n := len(nums)
	pre := 0
	start, end := n, 0
	minv := nums[n-1]
	for i := 1; i < n; i++ {
		if nums[pre] > nums[i] {
			start = min(pre, start)
			minv = min(minv, nums[i])
			end = i
		} else {
			pre = i
		}
	}
	if end == 0 {
		return 0
	}
	for i := 0; i < start; i++ {
		if nums[i] > minv {
			start = i
			break
		}
	}
	return max(0, end-start+1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
