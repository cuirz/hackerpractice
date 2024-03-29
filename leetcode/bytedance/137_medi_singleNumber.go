package main

//137. 只出现一次的数字 II
//给你一个整数数组nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
//
//
//
//示例 1：
//
//输入：nums = [2,2,3,2]
//输出：3
//示例 2：
//
//输入：nums = [0,1,0,1,0,1,99]
//输出：99
//
//
//提示：
//
//1 <= nums.length <= 3 * 10^4
//-2^31 <= nums[i] <= 2^31 - 1
//nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
//
//
//进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

//思路  位运算
func singleNumber(nums []int) int {
	var result int
	for i := 0; i < 32; i++ {
		sum := 0
		for _, v := range nums {
			sum += (v >> i) & 1
		}
		if sum%3 > 0 {
			result |= i << 1
		}
	}
	return result
}
