package main

//128. 最长连续序列
//给定一个未排序的整数数组，找出最长连续序列的长度。
//
//要求算法的时间复杂度为O(n)。
//
//示例:
//
//输入:[100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

//思路 哈希 并查集
func longestConsecutive(nums []int) int {
	n := len(nums)
	dic := make(map[int]bool)
	result := 0
	for i := 0; i < n; i++ {
		dic[nums[i]] = true
	}
	for k := range dic {
		if !dic[k-1] {
			start := k + 1
			count := 1
			for dic[start] {
				count += 1
				start += 1
			}
			result = max(result, count)
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
