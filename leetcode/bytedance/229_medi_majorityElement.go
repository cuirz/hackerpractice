package main

//229. 求众数 II
//给定一个大小为n的整数数组，找出其中所有出现超过⌊ n/3 ⌋次的元素。
//
//
//
//
//示例1：
//
//输入：[3,2,3]
//输出：[3]
//示例 2：
//
//输入：nums = [1]
//输出：[1]
//示例 3：
//
//输入：[1,1,1,3,3,2,2,2]
//输出：[1,2]
//
//提示：
//
//1 <= nums.length <= 5 * 10^4
//-10^9 <= nums[i] <= 10^9
//
//进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。

// 摩尔投票法
func majorityElement(nums []int) []int {
	elem1, elem2 := 0, 0
	vote1, vote2 := 0, 0
	for _, v := range nums {
		if vote1 > 0 && elem1 == v {
			vote1++
		} else if vote2 > 0 && elem2 == v {
			vote2++
		} else if vote1 == 0 {
			vote1++
			elem1 = v
		} else if vote2 == 0 {
			vote2++
			elem2 = v
		} else {
			vote1--
			vote2--
		}
	}
	cnt1, cnt2 := 0, 0
	for _, v := range nums {
		if vote1 > 0 && v == elem1 {
			cnt1++
		}
		if vote2 > 0 && v == elem2 {
			cnt2++
		}
	}

	var ans []int
	if vote1 > 0 && cnt1 > len(nums)/3 {
		ans = append(ans, elem1)
	}
	if vote2 > 0 && cnt2 > len(nums)/3 {
		ans = append(ans, elem2)
	}
	return ans
}
