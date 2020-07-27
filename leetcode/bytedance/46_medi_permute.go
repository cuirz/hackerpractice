package main

//46. 全排列
//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//[1,2,3],
//[1,3,2],
//[2,1,3],
//[2,3,1],
//[3,1,2],
//[3,2,1]
//]

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	n := len(nums)
	visited := make([]bool, n)
	var recursive func(array []int)
	recursive = func(array []int) {
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			dst := make([]int, len(array))
			copy(dst, array)
			dst = append(dst, nums[i])
			recursive(dst)
			visited[i] = false
		}
		if len(array) == n {
			result = append(result, array)
		}

	}
	recursive([]int{})
	return result
}
