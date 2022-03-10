package main

import "sort"

//215. 数组中的第K个最大元素
//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//示例 1:
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//示例2:
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//说明:
//
//你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

//思路 快速排序

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n-k]

}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := nums[0]
	head, tail, i := 0, len(nums)-1, 1
	for head < tail {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	nums[head] = mid
	quickSort(nums[:head])
	quickSort(nums[head+1:])
}
