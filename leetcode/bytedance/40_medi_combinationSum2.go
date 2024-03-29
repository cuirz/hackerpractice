package main

import (
	"sort"
)

//40. 组合总和 II
//给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
//
//candidates中的每个数字在每个组合中只能使用一次。
//
//说明：
//
//所有数字（包括目标数）都是正整数。
//解集不能包含重复的组合。
//示例1:
//
//输入: candidates =[10,1,2,7,6,1,5], target =8,
//所求解集为:
//[
//[1, 7],
//[1, 2, 5],
//[2, 6],
//[1, 1, 6]
//]
//示例2:
//
//输入: candidates =[2,5,2,1,2], target =5,
//所求解集为:
//[
// [1,2,2],
// [5]
//]

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var recursive func(nums []int, sum int) [][]int
	recursive = func(nums []int, sum int) [][]int {
		result := make([][]int, 0)
		for i, v := range nums {
			if i > 0 && v == nums[i-1] {
				continue
			} else if sum < v {
				break
			} else if sum == v {
				result = append(result, []int{v})
				break
			}
			res := recursive(nums[i+1:], sum-v)
			for _, e := range res {
				result = append(result, append([]int{v}, e...))
			}
		}
		return result
	}
	return recursive(candidates, target)
}

func main() {
	result := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	println(result)
}
