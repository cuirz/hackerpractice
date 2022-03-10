package main

import (
	"math"
	"sort"
)

//16. 最接近的三数之和
//给定一个包括n 个整数的数组nums和 一个目标值target。找出nums中的三个整数，使得它们的和与target最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
//
//示例：
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//提示：
//
//3 <= nums.length <= 10^3
//-10^3<= nums[i]<= 10^3
//-10^4<= target<= 10^4

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	two, sum := 0, 0
	closest := math.MaxInt32
	for i := 0; i < n; i++ {
		two = target - nums[i]
		l := i + 1
		r := n - 1
		top := i
		for l < r {
			sum = nums[l] + nums[r]
			if abs(two-sum) < abs(target-closest) {
				closest = nums[i] + sum
				if r-top == 1 {
					return closest
				}
			}
			if sum > two {
				r--
			} else if sum < two {
				l++
			} else {
				return target
			}
		}
	}
	return closest
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
