package main

import "sort"

//611. 有效三角形的个数
//给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
//
//示例 1:
//
//输入: [2,2,3,4]
//输出: 3
//解释:
//有效的组合是:
//2,3,4 (使用第一个 2)
//2,3,4 (使用第二个 2)
//2,2,3
//注意:
//
//数组长度不超过1000。
//数组里整数的范围为 [0, 1000]。

//思路  排序 + 二分查找
func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	var count int
	for i := 0; i < n-2; i++ {
		if nums[i] == 0 {
			continue
		}
		for j := i + 1; j < n-1; j++ {
			c := nums[i] + nums[j]
			index := sort.SearchInts(nums[j+1:], c)
			count += index
		}
	}
	return count
}

func main() {
	//println(sort.SearchInts([]int{2, 2, 2, 2}, 1))
	println(triangleNumber([]int{2,2,3,4}))
}
