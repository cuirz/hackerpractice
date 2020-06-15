package main

import "strconv"

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

//思路： 树,哈希表
type K []int

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	dic := make(map[string]int)

	result := make([][]int, 0)
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			v := append([]int{}, nums...)
			key := ""
			for k := 0; k < len(v); k++ {
				key += strconv.Itoa(v[k])
			}
			if dic[key] == 0 {
				dic[key] = 1
				result = append(result, v)
			}
		}
		for j := i; j < n; j++ {
			nums[i], nums[j] = nums[j], nums[i]
			backtrack(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	backtrack(0)
	return result
}

func main() {
	permuteUnique([]int{1, 1, 2})
}
