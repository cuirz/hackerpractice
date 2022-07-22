package main

import "sort"

//757. 设置交集大小至少为2
//一个整数区间[a, b](a < b) 代表着从a到b的所有连续整数，包括a和b。
//
//给你一组整数区间intervals，请找到一个最小的集合 S，使得 S 里的元素与区间intervals中的每一个整数区间都至少有2个元素相交。
//
//输出这个最小集合S的大小。
//
//示例 1:
//
//输入: intervals = [[1, 3], [1, 4], [2, 5], [3, 5]]
//输出: 3
//解释:
//考虑集合 S = {2, 3, 4}. S与intervals中的四个区间都有至少2个相交的元素。
//且这是S最小的情况，故我们输出3。
//示例 2:
//
//输入: intervals = [[1, 2], [2, 3], [2, 4], [4, 5]]
//输出: 5
//解释:
//最小的集合S = {1, 2, 3, 4, 5}.
//注意:
//
//intervals的长度范围为[1, 3000]。
//intervals[i]长度为2，分别代表左、右边界。
//intervals[i][j] 的值是[0, 10^8]范围内的整数

//贪婪
func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1])
	})
	n := len(intervals)
	result := 2
	left := intervals[n-1][0]
	right := intervals[n-1][0] + 1
	for i := n - 2; i >= 0; i-- {
		if intervals[i][1] >= right {
			continue
		} else if intervals[i][1] < left {
			left = intervals[i][0]
			right = intervals[i][0] + 1
			result += 2
		} else {
			right = left
			left = intervals[i][0]
			result += 1
		}
	}
	return result

}
