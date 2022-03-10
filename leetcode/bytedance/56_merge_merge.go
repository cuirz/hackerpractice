package main

import "sort"

//56. 合并区间
//给出一个区间的集合，请合并所有重叠的区间。
//
//示例 1:
//
//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例2:
//
//输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

//思路  排序 分治法

func merge(intervals [][]int) [][]int {

	result := make([][]int, 0)
	n := len(intervals)
	if n < 1 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	next := make([]int, 2)
	copy(next, intervals[0])
	for i := 0; i <= n; i++ {
		if i == n {
			result = append(result, next)
			break
		}
		if next[1] >= intervals[i][0] {
			next[1] = max(next[1], intervals[i][1])
		} else {
			result = append(result, next)
			next = make([]int, 2)
			copy(next, intervals[i])
		}
	}
	return result

}
