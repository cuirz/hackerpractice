package main

//41. 缺失的第一个正数
//给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
//
//
//示例 1:
//
//输入: [1,2,0]
//输出: 3
//示例 2:
//
//输入: [3,4,-1,1]
//输出: 2
//示例 3:
//
//输入: [7,8,9,11,12]
//输出: 1
//
//提示：
//
//你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。

//思路
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	var value int
	for i := 0; i < n; i++ {
		value = abs(nums[i])
		if value <= n {
			nums[value-1] = -abs(nums[value-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 思路： 原地哈希表
func firstMissingPositive2(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 1
	}
	var value int
	for i := 0; i < n; i++ {
		value = nums[i]
		if value < 1 {
			nums[i] = 0
			continue
		}
		for i+1 != value {
			if value > n || value < 1 || nums[value-1] == value {
				nums[i] = 0
				break
			}
			value, nums[value-1] = nums[value-1], value
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	println(firstMissingPositive([]int{1, 2}))
}
