package main

import "sort"

//315. 计算右侧小于当前元素的个数
//给定一个整数数组 nums，按要求返回一个新数组counts。数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于nums[i] 的元素的数量。
//
//示例:
//
//输入: [5,2,6,1]
//输出: [2,1,1,0]
//解释:
//5 的右侧有 2 个更小的元素 (2 和 1).
//2 的右侧仅有 1 个更小的元素 (1).
//6 的右侧有 1 个更小的元素 (1).
//1 的右侧有 0 个更小的元素.

//思路 排序， 二分查找
func countSmaller(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	heap := make([]int, n)
	copy(heap, nums)
	sort.Ints(heap)
	var search func(v, l, r int) int
	search = func(v, l, r int) int {
		mid := l + (r-l)>>1
		if heap[mid] == v {
			for mid > 0 && heap[mid-1] == v {
				mid -= 1
			}
			pre := heap[:mid]
			heap = append(pre, heap[mid+1:]...)
			return mid
		} else if heap[mid] < v {
			return search(v, mid+1, r)
		}
		return search(v, l, mid-1)

	}
	for i := 0; i < n; i++ {
		result[i] = search(nums[i], 0, len(heap)-1)
	}
	return result
}

func main() {
	println(countSmaller([]int{5, 2, 6, 1}))
}
