package main

//1893. 检查是否区域内所有整数都被覆盖
//给你一个二维整数数组ranges和两个整数left和right。每个ranges[i] = [starti, endi]表示一个从starti到endi的闭区间。
//
//如果闭区间[left, right]内每个整数都被ranges中至少一个区间覆盖，那么请你返回true，否则返回false。
//
//已知区间 ranges[i] = [starti, endi] ，如果整数 x 满足 starti <= x <= endi，那么我们称整数x被覆盖了。
//
//
//
//示例 1：
//
//输入：ranges = [[1,2],[3,4],[5,6]], left = 2, right = 5
//输出：true
//解释：2 到 5 的每个整数都被覆盖了：
//- 2 被第一个区间覆盖。
//- 3 和 4 被第二个区间覆盖。
//- 5 被第三个区间覆盖。
//示例 2：
//
//输入：ranges = [[1,10],[10,20]], left = 21, right = 21
//输出：false
//解释：21 没有被任何一个区间覆盖。
//
//
//提示：
//
//1 <= ranges.length <= 50
//1 <= starti <= endi <= 50
//1 <= left <= right <= 50

//思路 差分分组
func isCovered(ranges [][]int, left int, right int) bool {
	diff := make([]int, 52)
	for _, v := range ranges {
		diff[v[0]]++
		diff[v[1]+1]--
	}
	sum := 0
	for i := 1; i <= 50; i++ {
		sum += diff[i]
		for i >= left && i <= right && sum <= 0 {
			return false
		}
	}
	return true
}
