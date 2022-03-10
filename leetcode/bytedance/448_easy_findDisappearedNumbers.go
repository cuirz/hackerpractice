package main

//448. 找到所有数组中消失的数字
//给定一个范围在 1 ≤ a[i] ≤ n (n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
//
//找到所有在 [1, n] 范围之间没有出现在数组中的数字。
//
//您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
//
//示例:
//
//输入:
//[4,3,2,7,8,2,3,1]
//
//输出:
//[5,6]

//思路 哈希表

//思路 原地寻址
//不适用额外空间，
//寻找自己位置，找到了就换符号，未被换符号的部分就是消失的数字
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] *= -1
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findDisappearedNumbersHash(nums []int) []int {
	n := len(nums)
	dic := make(map[int]int, n)
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		dic[nums[i]] = 1
	}
	for i := 1; i <= n; i++ {
		if dic[i] == 0 {
			result = append(result, i)
		}
	}
	return result
}
