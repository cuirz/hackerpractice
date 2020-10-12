package main

import "sort"

//47. 全排列 II
//给定一个可包含重复数字的序列，返回所有不重复的全排列。
//
//示例:
//
//输入: [1,1,2]
//输出:
//[
//[1,1,2],
//[1,2,1],
//[2,1,1]
//]

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var dfs func(array []int) [][]int
	dfs = func(array []int) [][]int {
		var result [][]int
		if len(array) == 0 {
			return [][]int{array}
		}
		var before int
		for i, v := range array {
			if i == 0 || (i > 0 && v != before) {
				tmp := make([]int, i)
				copy(tmp, array[:i])
				tmp = append(tmp, array[i+1:]...)
				res := dfs(tmp)
				for j := 0; j < len(res); j++ {
					result = append(result, append(res[j], v))
				}
			}
			before = v
		}
		return result
	}
	return dfs(nums)
}
