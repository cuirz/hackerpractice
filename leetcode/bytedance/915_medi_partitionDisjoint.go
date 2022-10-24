package main

//915. 分割数组
//给定一个数组nums，将其划分为两个连续子数组left和right，使得：
//
//left中的每个元素都小于或等于right中的每个元素。
//left 和right都是非空的。
//left 的长度要尽可能小。
//在完成这样的分组后返回left的长度。
//
//用例可以保证存在这样的划分方法。
//
//
//
//示例 1：
//
//输入：nums = [5,0,3,8,6]
//输出：3
//解释：left = [5,0,3]，right = [8,6]
//示例 2：
//
//输入：nums = [1,1,1,0,6,12]
//输出：4
//解释：left = [1,1,1,0]，right = [6,12]
//
//提示：
//
//2 <= nums.length <= 10^5
//0 <= nums[i] <= 10^6
//可以保证至少有一种方法能够按题目所描述的那样对 nums 进行划分。

func partitionDisjoint(nums []int) int {
	left_max, cur_max := nums[0], nums[0]
	left_pos := 0
	for i := 1; i < len(nums)-1; i++ {
		cur_max = max(cur_max, nums[i])
		if nums[i] < left_max {
			left_max = cur_max
			left_pos = i
		}
	}
	return left_pos + 1
}
