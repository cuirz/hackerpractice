package main

//34. 在排序数组中查找元素的第一个和最后一个位置
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//你的算法时间复杂度必须是O(log n) 级别。
//
//如果数组中不存在目标值，返回[-1, -1]。
//
//示例 1:
//
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4]
//示例2:
//
//输入: nums = [5,7,7,8,8,10], target = 6
//输出: [-1,-1]

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n < 1 {
		return []int{-1, -1}
	}
	var search func(isLeft bool) int
	search = func(isLeft bool) int {
		l := 0
		// 特别注意 去长度不是下标
		r := len(nums)
		for l < r {
			mid := (l + r) >> 1
			if target < nums[mid] || (isLeft && target == nums[mid]) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return l
	}

	start := search(true)
	if start == n || nums[start] != target {
		return []int{-1, -1}
	}
	return []int{start, search(false) - 1}

}

func main() {
	println(searchRange([]int{1}, 1))
}
