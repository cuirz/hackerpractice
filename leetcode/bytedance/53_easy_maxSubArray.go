package main

//53. 最大子序和
//给定一个整数数组 nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释:连续子数组[4,-1,2,1] 的和最大，为6。
//进阶:
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

//思路 动态规划
//思路 分治法

//对于一个区间[l,r]，我们可以维护四个量：
//lSum 表示[l,r] 内以l 为左端点的最大子段和
//rSum 表示[l,r] 内以r 为右端点的最大子段和
//mSum 表示[l,r] 内的最大子段和
//iSum 表示[l,r] 的区间和

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	result := nums[0]
	sum := 0
	for _, v := range nums {
		sum += v
		result = max(result, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return result
}

func maxSubArrayBest(nums []int) int {
	type Status struct {
		iSum, lSum, rSum, mSum int
	}
	var pageUp func(l, r Status) Status
	var get func(l, r int) Status
	pageUp = func(l, r Status) Status {
		isum := l.iSum + r.iSum
		lsum := max(l.lSum, l.iSum+r.lSum)
		rsum := max(r.rSum, r.iSum+l.rSum)
		msum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
		return Status{iSum: isum, lSum: lsum, rSum: rsum, mSum: msum}
	}
	get = func(l, r int) Status {
		if l == r {
			return Status{nums[l], nums[l], nums[l], nums[l]}
		}
		mid := (l + r) >> 1
		left := get(l, mid)
		right := get(mid+1, r)
		return pageUp(left, right)
	}
	return get(0, len(nums)-1).mSum
}
