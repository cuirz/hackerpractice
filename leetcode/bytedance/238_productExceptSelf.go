package main

//238. 除自身以外数组的乘积
//给你一个长度为n的整数数组nums，其中n > 1，返回输出数组output，其中 output[i]等于nums中除nums[i]之外其余各元素的乘积。
//
//输入: [1,2,3,4]
//输出: [24,12,8,6]

//提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
//
//说明: 请不要使用除法，且在O(n) 时间复杂度内完成此题。
//
//进阶：
//你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

func productExceptSelf(nums []int) []int {

	n := len(nums)
	result := make([]int, n)
	pre := make([]int, n)
	suf := make([]int, n)
	pre[0], suf[0] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i]
		suf[i] = suf[i-1] * nums[n-i-1]
	}
	for i := 1; i < n-1; i++ {
		result[i] = pre[i-1] * suf[n-i-2]
	}
	result[0] = suf[n-2]
	result[n-1] = pre[n-2]
	return result
}

func productExceptSelfO1(nums []int) []int {

	n := len(nums)
	result := make([]int, n)
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	right := 1
	for i := n - 1; i > -1; i-- {
		result[i] *= right
		right *= nums[i]
	}
	return result
}
