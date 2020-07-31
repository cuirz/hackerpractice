package main

//78. 子集
//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入: nums = [1,2,3]
//输出:
//[
//[3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]

//思路 递归
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	for _, v := range nums {
		for _, k := range result {
			sub := []int{v}
			sub = append(sub, k...)
			result = append(result, sub)
		}
		result = append(result, []int{v})
	}
	result = append(result,[]int{})
	return result

}
