package main

//给定一个无序的整数数组，找到其中最长上升子序列的长度。
//
//示例:
//
//输入: [10,9,2,5,3,7,101,18]
//输出: 4
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
//说明:
//
//可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
//你算法的时间复杂度应该为 O(n2) 。
//进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

//思路： 动态规划
//思路： 贪心 + 二分查找

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	length := 1
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] > dp[length-1] {
			dp[length] = nums[i]
			length++
		} else {
			l, r = 0, length-1
			for l < r {
				mid := (l + r - 1) / 2
				if dp[mid] < nums[i] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			dp[l] = nums[i]
		}
	}

	return length
}

func main() {
	println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
