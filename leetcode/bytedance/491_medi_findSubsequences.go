package main

import "math"

//491. 递增子序列
//给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
//
//示例:
//
//输入: [4, 6, 7, 7]
//输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
//说明:
//
//给定数组的长度不会超过15。
//数组中的整数范围是[-100,100]。
//给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。

//思路  深度优先
// 对数组中 当前下标 选择和不选择 进行深度递归
func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	array := make([]int, 0)
	var dfs func(index int, pre int)
	dfs = func(index int, pre int) {
		if index == len(nums) {
			if len(array) >= 2 {
				tmp := make([]int, len(array))
				copy(tmp, array)
				result = append(result, tmp)
			}
			return
		}
		if nums[index] >= pre {
			array = append(array, nums[index])
			dfs(index+1, nums[index])
			array = array[:len(array)-1]
		}
		if nums[index] != pre {
			dfs(index+1, pre)
		}
	}
	dfs(0, math.MinInt32)
	return result
}

func main() {
	println(findSubsequences([]int{4, 6, 7, 7}))
}
