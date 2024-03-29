package main

//57. 插入区间
//给出一个无重叠的 ，按照区间起始端点排序的区间列表。
//
//在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
//
//
//示例1：
//
//输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
//输出：[[1,5],[6,9]]
//示例2：
//
//输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出：[[1,2],[3,10],[12,16]]
//解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10]重叠。
//
//
//注意：输入类型已在 2019 年 4 月 15 日更改。请重置为默认代码定义以获取新的方法签名。

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	for i, v := range intervals {
		if newInterval[1] < v[0] {
			result = append(result, newInterval, v)
			result = append(result, intervals[i+1:]...)
			return result
		} else if v[1] < newInterval[0] {
			result = append(result, v)
		} else {
			newInterval[0] = min(v[0], newInterval[0])
			newInterval[1] = max(v[1], newInterval[1])
		}
	}
	result = append(result, newInterval)
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
