package main

import "math"

//209. 长度最小的子数组
//给定一个含有n个正整数的数组和一个正整数s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。
//
//示例:
//
//输入: s = 7, nums = [2,3,1,2,4,3]
//输出: 2
//解释: 子数组[4,3]是该条件下的长度最小的连续子数组。
//进阶:
//
//如果你已经完成了O(n) 时间复杂度的解法, 请尝试O(n log n) 时间复杂度的解法。

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)

	result, sum, start := math.MaxInt32, 0, 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		if sum >= s {
			result = min(result, i-start+1)
			for start < i {
				sum -= nums[start]
				start++
				if sum < s {
					break
				}
				result = min(result, i-start+1)
			}
		}
	}
	if result == math.MaxInt32 {
		return 0
	}

	return result
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))

}
