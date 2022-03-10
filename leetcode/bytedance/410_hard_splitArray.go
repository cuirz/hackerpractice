package main

import "math"

//410. 分割数组的最大值
//给定一个非负整数数组和一个整数m，你需要将这个数组分成m个非空的连续子数组。设计一个算法使得这m个子数组各自和的最大值最小。
//
//注意:
//数组长度n满足以下条件:
//
//1 ≤ n ≤ 1000
//1 ≤ m ≤ min(50, n)
//示例:
//
//输入:
//nums = [7,2,5,10,8]
//m = 2
//
//输出:
//18
//
//解释:
//一共有四种方法将nums分割为2个子数组。
//其中最好的方式是将其分为[7,2,5] 和 [10,8]，
//因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。

//思路 动态规划
// dp[i][j]  数组中 i个被分割成j段中，子数组中和为最小值
// dp[i][j] = min(dp[i][j],max(dp[k][j-1],sub(k+1,i)))

//思路 二分查找
// 数组nums中能 取到的子数组和范围，[nums中最大值，nums数组和]
// 把这个子数组范围抽象出来进行二分查找，既匹配m段又匹配子数组中"最大"最小和

// 【7，2，5，10，8，9，1，15，17，4，6】

func splitArray(nums []int, m int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	sub := make([]int, n+1)
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	for i := 0; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], sub[i]-sub[k]))
			}
		}
	}
	return dp[n][m]
}

// 二分查找
func splitArray(nums []int, m int) int {
	n := len(nums)
	left, right := 0, 0
	for i := 0; i < n; i++ {
		right += nums[i]
		if left < nums[i] {
			left = nums[i]
		}
	}
	check := func(mid int) bool {
		sum := 0
		count := 1
		for i := 0; i < n; i++ {
			if sum+nums[i] > mid {
				count++
				sum = nums[i]
			} else {
				sum += nums[i]
			}
		}
		return count <= m
	}

	for left < right {
		mid := (left + right) >> 1
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
