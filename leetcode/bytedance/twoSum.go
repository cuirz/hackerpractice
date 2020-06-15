package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{-1, -1}
	}
	for first := 0; first < len(nums); first++ {
		for second := first + 1; second < len(nums); second++ {
			if target == nums[first]+nums[second] {
				return []int{first, second}
			}
		}
	}
	return []int{-1, -1}
}

func twoSum(nums []int, target int) []int {
	dic := map[int]int{}
	for i := 0; i < len(nums); i++ {
		dic[nums[i]] = i
	}
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	sum := 0
	result := make([]int, 0)
	for l < r {
		sum = nums[l] + nums[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			result = append(result, dic[nums[l]])
			result = append(result, dic[nums[r]])
			l++
			r--
		}
	}
	return result
}

func main() {
	fmt.Println("%v", twoSum([]int{3, 2, 3}, 6))
}
