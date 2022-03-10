package main

//665. 非递减数列
//给你一个长度为n的整数数组，请你判断在 最多 改变1 个元素的情况下，该数组能否变成一个非递减数列。
//
//我们是这样定义一个非递减数列的：对于数组中所有的i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。
//
//
//
//示例 1:
//
//输入: nums = [4,2,3]
//输出: true
//解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
//示例 2:
//
//输入: nums = [4,2,1]
//输出: false
//解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
//
//
//说明：
//
//1 <= n <= 10 ^ 4
//- 10 ^ 5<= nums[i] <= 10 ^ 5

//思路 峰值算法
func checkPossibility(nums []int) bool {
	n := len(nums)
	count := 0
	for i := 0; i < n-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			count++
			if count > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}
