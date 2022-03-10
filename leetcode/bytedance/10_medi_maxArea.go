package main

//11. 盛最多水的容器
//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0)。找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器，且n的值至少为 2。
//示例：
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49

// 思路 滑动窗口，双指针

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	result := 0
	for l < r {
		result = max(result, (r-l)*min(height[l], height[r]))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
