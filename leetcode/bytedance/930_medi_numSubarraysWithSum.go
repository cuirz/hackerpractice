package main

//930. 和相同的二元子数组
//给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。
//
//子数组 是数组的一段连续部分。
//
//
//示例 1：
//
//输入：nums = [1,0,1,0,1], goal = 2
//输出：4
//解释：
//如下面黑体所示，有 4 个满足题目要求的子数组：
//[1,0,1,0,1]
//[1,0,1,0,1]
//[1,0,1,0,1]
//[1,0,1,0,1]
//示例 2：
//
//输入：nums = [0,0,0,0,0], goal = 0
//输出：15
//
//提示：
//
//1 <= nums.length <= 3 * 10^4
//nums[i] 不是 0 就是 1
//0 <= goal <= nums.length

//思路 滑动窗口 哈希表
func numSubarraysWithSum(nums []int, goal int) int {
	count := make(map[int]int)
	sum := 0
	var result int
	for _,v := range nums{
		count[sum]++
		sum += v
		result += count[sum-goal]
	}
	return result

}
func main() {
	println(numSubarraysWithSum([]int{0,1,1,1,1}, 3))
}
