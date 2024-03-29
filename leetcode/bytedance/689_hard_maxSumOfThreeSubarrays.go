package main

//689. 三个无重叠子数组的最大和
//给你一个整数数组 nums 和一个整数 k ，找出三个长度为 k 、互不重叠、且3 * k 项的和最大的子数组，并返回这三个子数组。
//
//以下标的数组形式返回结果，数组中的每一项分别指示每个子数组的起始位置（下标从 0 开始）。如果有多个结果，返回字典序最小的一个。
//
//
//
//示例 1：
//
//输入：nums = [1,2,1,2,6,7,5,1], k = 2
//输出：[0,3,5]
//解释：子数组 [1, 2], [2, 6], [7, 5] 对应的起始下标为 [0, 3, 5]。
//也可以取 [2, 1], 但是结果 [1, 3, 5] 在字典序上更大。
//示例 2：
//
//输入：nums = [1,2,1,2,1,2,1,2,1], k = 2
//输出：[0,2,4]
//
//
//提示：
//
//1 <= nums.length <= 2 * 10^4
//1 <= nums[i] <2^16
//1 <= k <= floor(nums.length / 3)

//思路 滑动窗口

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	index1, maxIndex1, index2 := 0, 0, 0
	sum1, sum2, sum3 := 0, 0, 0
	maxSum1, maxSum2, maxSum3 := 0, 0, 0
	ans := make([]int, 0)
	for i := k * 2; i < n; i++ {
		sum1 += nums[i-k*2]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= k*3-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				index1 = i - k*3 + 1
			}
			if sum2+maxSum1 > maxSum2 {
				maxSum2 = sum2 + maxSum1
				maxIndex1, index2 = index1, i-k*2+1
			}
			if sum3+maxSum2 > maxSum3 {
				maxSum3 = sum3 + maxSum2
				ans = []int{maxIndex1, index2, i - k + 1}
			}
			sum1 -= nums[i-k*3+1]
			sum2 -= nums[i-k*2+1]
			sum3 -= nums[i-k+1]
		}
	}
	return ans
}

func main() {
	println(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 6, 7, 5, 1}, 2))
}
