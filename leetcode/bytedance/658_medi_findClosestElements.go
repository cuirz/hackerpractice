package main

import "sort"

//658. 找到 K 个最接近的元素
//给定一个 排序好 的数组arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
//
//整数 a 比整数 b 更接近 x 需要满足：
//
//|a - x| < |b - x| 或者
//|a - x| == |b - x| 且 a < b
//
//
//示例 1：
//
//输入：arr = [1,2,3,4,5], k = 4, x = 3
//输出：[1,2,3,4]
//示例 2：
//
//输入：arr = [1,2,3,4,5], k = 4, x = -1
//输出：[1,2,3,4]
//
//
//提示：
//
//1 <= k <= arr.length
//1 <= arr.length<= 10^4
//arr按 升序 排列
//-10^4<= arr[i], x <= 10^4

func findClosestElements(arr []int, k int, x int) []int {
	right := sort.SearchInts(arr, x)
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= len(arr) || x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]

}
