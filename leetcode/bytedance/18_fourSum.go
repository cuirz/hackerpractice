package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	if size < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	l, r := 0, 0
	var result [][]int
	sum := 0
	for i := 0; i < size-3; i++ {
		// 排重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < size-2; j++ {
			// 排重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l = j + 1
			r = size - 1
			sum = target - (nums[i] + nums[j])
			// 左右指针遍历
			for l < r {
				two := nums[l] + nums[r]
				if two == sum {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if two < sum {
					l++
				} else {
					r--
				}
			}
		}
	}
	return result
}

func main() {
	println(fmt.Sprintf("%v", fourSum([]int{1, 0, -1, 0, -2, 2}, 0)))
}
