package main

//33. 搜索旋转排序数组
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组[0,1,2,4,5,6,7]可能变为[4,5,6,7,0,1,2])。
//
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回-1。
//
//你可以假设数组中不存在重复的元素。
//
//你的算法时间复杂度必须是O(logn) 级别。
//
//示例 1:
//
//输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
//示例2:
//
//输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1

//思路  二分查找

func search(nums []int, target int) int {
	var binarySearch func(array []int, leftIndex int) int
	binarySearch = func(array []int, leftIndex int) int {
		n := len(array)
		if n < 1 {
			return -1
		} else if n == 1 {
			if array[0] == target {
				return leftIndex
			}
			return -1
		}
		mid := n >> 1
		if (array[mid] > array[0] && target >= array[0] && target < array[mid]) ||
			(array[mid] < array[0] && (target >= array[0] || target < array[mid])) {
			return binarySearch(array[:mid], leftIndex)
		}
		return binarySearch(array[mid:], leftIndex+mid)
	}
	return binarySearch(nums, 0)
}

func main() {
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
