package main

//805. 数组的均值分割
//给定你一个整数数组 nums
//
//我们要将 nums 数组中的每个元素移动到 A 数组 或者 B 数组中，使得 A 数组和 B 数组不为空，并且 average(A) == average(B) 。
//
//如果可以完成则返回true ， 否则返回 false  。
//
//注意：对于数组 arr ,  average(arr) 是 arr 的所有元素除以 arr 长度的和。
//
//
//
//示例 1:
//
//输入: nums = [1,2,3,4,5,6,7,8]
//输出: true
//解释: 我们可以将数组分割为 [1,4,5,8] 和 [2,3,6,7], 他们的平均值都是4.5。
//示例 2:
//
//输入: nums = [3,1]
//输出: false
//
//
//提示:
//
//1 <= nums.length <= 30
//0 <= nums[i] <= 10^4
func splitArraySameAverage(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	n := len(nums)
	m := n / 2
	isPossible := false
	for i := 1; i <= m; i++ {
		if sum*i%n == 0 {
			isPossible = true
			break
		}
	}
	if !isPossible {
		return false
	}

	dp := make([]map[int]bool, m+1)
	for i := range dp {
		dp[i] = map[int]bool{}
	}
	dp[0][0] = true
	for _, num := range nums {
		for i := m; i >= 1; i-- {
			for x := range dp[i-1] {
				curr := x + num
				if curr*n == sum*i {
					return true
				}
				dp[i][curr] = true
			}
		}
	}
	return false
}
