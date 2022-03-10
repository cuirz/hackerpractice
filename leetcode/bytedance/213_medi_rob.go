package main

//213. 打家劫舍 II
//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
//示例1:
//
//输入: [2,3,2]
//输出: 3
//解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//示例 2:
//
//输入: [1,2,3,1]
//输出: 4
//解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//    偷窃到的最高金额 = 1 + 3 = 4 。

func rob(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	pre, res := 0, 0
	pre1, res1 := 0, 0

	for i, v := range nums {
		if i < n-1 {
			pre, res = res, max(pre+v, res)
		}
		if i > 0 {
			pre1, res1 = res1, max(pre1+v, res1)
		}

	}
	return max(res, res1)

}
