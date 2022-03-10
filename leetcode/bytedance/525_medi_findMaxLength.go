package main

//525. 连续数组
//给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
//
//
//示例 1:
//
//输入: nums = [0,1]
//输出: 2
//说明: [0, 1] 是具有相同数量0和1的最长连续子数组。
//示例 2:
//
//输入: nums = [0,1,0]
//输出: 2
//说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
//
//
//提示：
//
//1 <= nums.length <= 10^5
//nums[i] 不是 0 就是 1

//思路 前缀和 ， 哈希表
func findMaxLength(nums []int) int {
	count := 0
	dic := map[int]int{0: -1}
	var result int
	for i, v := range nums {
		if v == 0 {
			count--
		} else {
			count++
		}
		if index, has := dic[count]; has {
			result = max(result, i-index)
		} else {
			dic[count] = i
		}
	}
	return result
}

func main() {
	println(findMaxLength([]int{0, 1, 0, 0, 1}))
}
