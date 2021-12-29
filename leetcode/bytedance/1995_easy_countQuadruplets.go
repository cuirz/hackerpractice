package main

//1995. 统计特殊四元组
//给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：
//
//nums[a] + nums[b] + nums[c] == nums[d] ，且
//a < b < c < d
//
//
//示例 1：
//
//输入：nums = [1,2,3,6]
//输出：1
//解释：满足要求的唯一一个四元组是 (0, 1, 2, 3) 因为 1 + 2 + 3 == 6 。
//示例 2：
//
//输入：nums = [3,3,6,4,5]
//输出：0
//解释：[3,3,6,4,5] 中不存在满足要求的四元组。
//示例 3：
//
//输入：nums = [1,1,1,3,5]
//输出：4
//解释：满足要求的 4 个四元组如下：
//- (0, 1, 2, 3): 1 + 1 + 1 == 3
//- (0, 1, 3, 4): 1 + 1 + 3 == 5
//- (0, 2, 3, 4): 1 + 1 + 3 == 5
//- (1, 2, 3, 4): 1 + 1 + 3 == 5
//
//
//提示：
//
//4 <= nums.length <= 50
//1 <= nums[i] <= 100

//思路  nums[a]+nums[b] = nums[d]-nums[c]
// c 取值 [b+1,n-1]
// b 取值 [1,n-3]
func countQuadruplets(nums []int) int {
	n := len(nums)
	dic := make(map[int]int)
	var ans int
	for b := n - 3; b >= 1; b-- {
		for d := b + 2; d < n; d++ {
			dic[nums[d]-nums[b+1]]++
		}
		for a := 0; a < b; a++ {
			ans += dic[nums[a]+nums[b]]
		}
	}
	return ans
}
